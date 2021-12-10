begin;

  -- wh_private_address_status returns a warehouse appropriate string
  -- representing if the address is private, public, or not applicable for the
  -- provided address.  An address which cannot be cast to an inet results in
  -- 'Not Applicable' being returned.
  create function wh_private_address_status(inet) returns text
  as $$
  begin
    case
      when $1 << any ('{10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12, fc00::/7, fe80::/10}'::cidr[]) then
        return 'Private';
      else
        return 'Public';
      end case;
  end;
  $$ language plpgsql;

  create view whx_network_address_dimension_source as
    select
      hdns.host_id as host_id,
      hdns.name as address,
      'DNS Name' as address_type,
      'Not Applicable' as ip_address_family,
      'Not Applicable' private_ip_address_status,
      hdns.name as dns_name,
      'Not Applicable' as ip4_address,
      'Not Applicable' as ip6_address
    from host_dns_name as hdns
    union
    select -- id is the first column in the target view
       hip.host_id as host_id,
       host(hip.address) as address,
       'IP Address' as address_type,
       case
         when family(hip.address) = 4 then 'IPv4'
         when family(hip.address) = 6 then 'IPv6'
         else 'Not Applicable'
         end               as ip_address_family,
       wh_private_address_status(hip.address) as private_ip_address_status,
       'Not Applicable' as dns_name,
       case
         when hip.address is not null and family(hip.address) = 4 then host(hip.address)
         else 'Not Applicable'
         end as ip4_address,
       case
         when hip.address is not null and family(hip.address) = 6 then host(hip.address)
         else 'Not Applicable'
         end as ip6_address
    from host_ip_address as hip;

commit;
