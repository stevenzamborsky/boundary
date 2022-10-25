begin;

  alter table session add column certificate_key_id text;

  update session set certificate_key_id = key_id;

  alter table session
    alter column certificate_key_id set not null
    add constraint kms_data_key_version_fkey foreign key (certificate_key_id) references kms_data_key_version (private_id);

  update session set key_id = (
    select kdkv.private_key
    from kms_data_key_version kdkv
    inner join kms_data_key   kdk
      on kdkv.data_key_id = kdk.private_id
    inner join kms_root_key   krk
      on kdk.root_key_id = krk.private_id
    where krk.scope_id = session.scope_id
      and kdk.purpose = 'database'
  )
  where tofu_token != null;

commit;
