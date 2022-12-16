begin;

-- Update table
alter table target_tcp
  add column default_client_port int; -- default_client_port can be null

-- Update views
-- Replaces target_all_subtypes defined in 59/01_target_ingress_egress_worker_filters.up.sql
drop view target_all_subtypes;
create view target_all_subtypes as
select public_id,
   project_id,
   name,
   description,
   default_port,
   default_client_port,
   session_max_seconds,
   session_connection_limit,
   version,
   create_time,
   update_time,
   worker_filter,
   egress_worker_filter,
   ingress_worker_filter,
   'tcp' as type
from target_tcp;

commit;
