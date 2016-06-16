# webhyve
A web interface for (vm-)bhyve

----
Currently fairly messy code, using vm-bhyve as the base to grab info. 
---

TODO: 
1. Pull all the VMs into a SQL(ite) database. 
2. Be able to start/stop VMs. 

initial planned handlers:

/view - json output of all vms
/start/name/$name - start VM by name
/start/uuid/$uuid - start VM by uuid
/stop/name/$name - stop VM by name
/stop/uuid/$uuid - stop VM by uuid

Creating/deleting VMs - somewhere much further down the line. 
