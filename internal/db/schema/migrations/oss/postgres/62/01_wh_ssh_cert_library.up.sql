begin;
  alter table wh_credential_dimension
    add column credential_library_username         text,
    add column credential_library_key_type         text,
    add column credential_library_key_bits         integer,
    add column credential_library_tll              text,
    add column credential_library_key_id           text,
    add column credential_library_critical_options text,
    add column credential_library_extensions       text
  ;

  update wh_credential_dimension set
    credential_library_username         = 'Not Applicable',
    credential_library_key_type         = 'Not Applicable',
    credential_library_key_bits         = -1,
    credential_library_tll              = 'Not Applicable',
    credential_library_key_id           = 'Not Applicable',
    credential_library_critical_options = 'Not Applicable',
    credential_library_extensions       = 'Not Applicable';

  alter table wh_credential_dimension
    alter column credential_library_username         type wh_dim_text,
    alter column credential_library_key_type         type wh_dim_text,
    alter column credential_library_key_bits          set not null,
    alter column credential_library_tll              type wh_dim_text,
    alter column credential_library_key_id           type wh_dim_text,
    alter column credential_library_critical_options type wh_dim_text,
    alter column credential_library_extensions       type wh_dim_text
  ;

  update wh_credential_dimension set
    credential_library_type = 'vault generic credential library'
  where
    credential_library_type = 'vault credential library';

  -- The whx_credential_dimension_source view shows the current values in the
  -- operational tables of the credential dimension.
  -- Replaces whx_credential_dimension_source defined in 44/03_targets.up.sql
  drop view whx_credential_dimension_source;
  create view whx_credential_dimension_source as
       select -- id is the first column in the target view
              s.public_id                              as session_id,
              coalesce(scd.credential_purpose, 'None') as credential_purpose,
              cl.public_id                             as credential_library_id,
              case
                when   vcl.public_id is not null then 'vault generic credential library'
                when vsccl.public_id is not null then 'vault ssh certificate credential library'
                else 'None'
                end                                    as credential_library_type,
              case
                when   vcl.public_id is not null then coalesce(  vcl.name, 'None')
                when vsccl.public_id is not null then coalesce(vsccl.name, 'None')
                else 'Unknown'
              end                                      as credential_library_name,
              case
                when   vcl.public_id is not null then coalesce(  vcl.description, 'None')
                when vsccl.public_id is not null then coalesce(vsccl.description, 'None')
                else 'Unknown'
              end                                      as credential_library_description,
              case
                when   vcl.public_id is not null then coalesce(  vcl.vault_path, 'None')
                when vsccl.public_id is not null then coalesce(vsccl.vault_path, 'None')
                else 'Unknown'
              end                                      as credential_library_vault_path,
              case
                when   vcl.public_id is not null then coalesce(vcl.http_method, 'None')
                when vsccl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_vault_http_method,
              case
                when   vcl.public_id is not null then coalesce(vcl.http_request_body::text, 'None')
                when vsccl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_vault_http_request_body,
              case
                when vsccl.public_id is not null then coalesce(vsccl.username, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_username,
              case
                when vsccl.public_id is not null then coalesce(vsccl.key_type, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_key_type,
              case
                when vsccl.public_id is not null then vsccl.key_bits
                when   vcl.public_id is not null then -1
                else -1
              end                                      as credential_library_key_bits,
              case
                when vsccl.public_id is not null then coalesce(vsccl.ttl, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_tll,
              case
                when vsccl.public_id is not null then coalesce(vsccl.key_id, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_key_id,
              case
                when vsccl.public_id is not null then coalesce(vsccl.critical_options::text, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_critical_options,
              case
                when vsccl.public_id is not null then coalesce(vsccl.extensions::text, 'None')
                when   vcl.public_id is not null then 'Not Applicable'
                else 'Unknown'
              end                                      as credential_library_extensions,
              cs.public_id                             as credential_store_id,
              case
                when vcs is null then 'None'
                else 'vault credential store'
                end                                    as credential_store_type,
              coalesce(vcs.name, 'None')               as credential_store_name,
              coalesce(vcs.description, 'None')        as credential_store_description,
              coalesce(vcs.namespace, 'None')          as credential_store_vault_namespace,
              coalesce(vcs.vault_address, 'None')      as credential_store_vault_address,
              t.public_id                              as target_id,
              'tcp target'                             as target_type,
              coalesce(tt.name, 'None')                as target_name,
              coalesce(tt.description, 'None')         as target_description,
              coalesce(tt.default_port, 0)             as target_default_port_number,
              tt.session_max_seconds                   as target_session_max_seconds,
              tt.session_connection_limit              as target_session_connection_limit,
              p.public_id                              as project_id,
              coalesce(p.name, 'None')                 as project_name,
              coalesce(p.description, 'None')          as project_description,
              o.public_id                              as organization_id,
              coalesce(o.name, 'None')                 as organization_name,
              coalesce(o.description, 'None')          as organization_description
      from session_credential_dynamic as scd
        join session            as s  on scd.session_id = s.public_id
        join credential_library as cl on scd.library_id = cl.public_id
        join credential_store   as cs on cl.store_id    = cs.public_id
        join target             as t  on s.target_id    = t.public_id
        join iam_scope          as p  on p.public_id    = t.project_id and p.type = 'project'
        join iam_scope          as o  on p.parent_id    = o.public_id  and o.type = 'org'

        left join credential_vault_library          as vcl   on cl.public_id = vcl.public_id
        left join credential_vault_ssh_cert_library as vsccl on cl.public_id = vsccl.public_id
        left join credential_vault_store            as vcs   on cs.public_id = vcs.public_id
        left join target_tcp                        as tt    on t.public_id  = tt.public_id;

  -- Replaces view in 16/02_wh_credential_dimension.up.sql
  drop view whx_credential_dimension_target;
  create view whx_credential_dimension_target as
  select key,
         credential_purpose,
         credential_library_id,
         credential_library_type,
         credential_library_name,
         credential_library_description,
         credential_library_vault_path,
         credential_library_vault_http_method,
         credential_library_vault_http_request_body,
         credential_library_username,
         credential_library_key_type,
         credential_library_key_bits,
         credential_library_tll,
         credential_library_key_id,
         credential_library_critical_options,
         credential_library_extensions,
         credential_store_id,
         credential_store_type,
         credential_store_name,
         credential_store_description,
         credential_store_vault_namespace,
         credential_store_vault_address,
         target_id,
         target_type,
         target_name,
         target_description,
         target_default_port_number,
         target_session_max_seconds,
         target_session_connection_limit,
         project_id,
         project_name,
         project_description,
         organization_id,
         organization_name,
         organization_description
    from wh_credential_dimension
   where current_row_indicator = 'Current'
  ;

  -- Replaces function in 16/03_wh_credential_dimension.up.sql
  drop function wh_upsert_credential_dimension;
  create function wh_upsert_credential_dimension(p_session_id wt_public_id, p_library_id wt_public_id, p_credential_purpose wh_dim_text) returns wh_dim_key
  as $$
  declare
    src     whx_credential_dimension_target%rowtype;
    target  whx_credential_dimension_target%rowtype;
    new_row wh_credential_dimension%rowtype;
    t_id    wt_public_id;
  begin
    select s.target_id into strict t_id
      from session as s
     where s.public_id = p_session_id;

    select * into target
      from whx_credential_dimension_target as t
     where t.credential_library_id = p_library_id
       and t.target_id             = t_id
       and t.credential_purpose    = p_credential_purpose;

    select
      target.key,                    t.credential_purpose,
      t.credential_library_id,       t.credential_library_type,     t.credential_library_name,     t.credential_library_description, t.credential_library_vault_path,    t.credential_library_vault_http_method, t.credential_library_vault_http_request_body,
      t.credential_library_username, t.credential_library_key_type, t.credential_library_key_bits, t.credential_library_tll,         t.credential_library_key_id,        t.credential_library_critical_options,  t.credential_library_extensions,
      t.credential_store_id,         t.credential_store_type,       t.credential_store_name,       t.credential_store_description,   t.credential_store_vault_namespace, t.credential_store_vault_address,
      t.target_id,                   t.target_type,                 t.target_name,                 t.target_description,             t.target_default_port_number,       t.target_session_max_seconds,           t.target_session_connection_limit,
      t.project_id,                  t.project_name,                t.project_description,
      t.organization_id,             t.organization_name,           t.organization_description
      into src
      from whx_credential_dimension_source as t
     where t.credential_library_id = p_library_id
       and t.session_id            = p_session_id
       and t.target_id             = t_id
       and t.credential_purpose    = p_credential_purpose;

    if src is distinct from target then
      update wh_credential_dimension
         set current_row_indicator = 'Expired',
             row_expiration_time   = current_timestamp
       where credential_library_id = p_library_id
         and target_id             = t_id
         and credential_purpose    = p_credential_purpose
         and current_row_indicator = 'Current';

      insert into wh_credential_dimension (
             credential_purpose,
             credential_library_id,       credential_library_type,     credential_library_name,     credential_library_description, credential_library_vault_path,    credential_library_vault_http_method, credential_library_vault_http_request_body,
             credential_library_username, credential_library_key_type, credential_library_key_bits, credential_library_tll,         credential_library_key_id,        credential_library_critical_options,  credential_library_extensions,
             credential_store_id,         credential_store_type,       credential_store_name,       credential_store_description,   credential_store_vault_namespace, credential_store_vault_address,
             target_id,                   target_type,                 target_name,                 target_description,             target_default_port_number,       target_session_max_seconds,           target_session_connection_limit,
             project_id,                  project_name,                project_description,
             organization_id,             organization_name,           organization_description,
             current_row_indicator,       row_effective_time,          row_expiration_time
      )
      select credential_purpose,
             credential_library_id,       credential_library_type,     credential_library_name,     credential_library_description, credential_library_vault_path,    credential_library_vault_http_method, credential_library_vault_http_request_body,
             credential_library_username, credential_library_key_type, credential_library_key_bits, credential_library_tll,         credential_library_key_id,        credential_library_critical_options,  credential_library_extensions,
             credential_store_id,         credential_store_type,       credential_store_name,       credential_store_description,   credential_store_vault_namespace, credential_store_vault_address,
             target_id,                   target_type,                 target_name,                 target_description,             target_default_port_number,       target_session_max_seconds,           target_session_connection_limit,
             project_id,                  project_name,                project_description,
             organization_id,             organization_name,           organization_description,
             'Current',                   current_timestamp,           'infinity'::timestamptz
        from whx_credential_dimension_source
       where credential_library_id = p_library_id
         and session_id            = p_session_id
         and target_id             = t_id
         and credential_purpose    = p_credential_purpose
      returning * into new_row;

      return new_row.key;
    end if;

    return target.key;
  end
  $$ language plpgsql;
commit;
