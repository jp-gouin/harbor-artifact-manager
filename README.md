# Harbor offline assets management

This software offer a management of private artifacts hosted on Harbor

It leverage the team to create project , add members and artifacts.

## Definition

An **artifact** is a combination of a Helm chart and all necessary docker images and dependencies.

## Feature
* Create project
* Add owner/member to a project
* Shopping view for adding artifact to project
* Build a datapack for air-gap deployment
   * The datapack if composed of helm charts
   * Docker images bundled by application (/w dependencies)
   * Custom Git resources (Quickstart for instance or Example) 
   * Script to unload datapack to an offline Harbor (or other docker registry)
* Add validation on artifact 
* Browse all artifacts and all versions available in the registry

## Dependencies
This solution requires : 
* ETCD
* Jenkins
* GCS Account (optional)
* LDAP server (optional)

## Installation
1. Clone the repo
2. Run 
```
docker build -t harbor-mgt-artifact .
```
3. create an env var file
```
  LISTEN=8080
  HARBOR=https://<harbor_url>
  USER=<harbor_user>
  PASWORD=<harbor_password>
  ETCD=192.168.65.1:2379
  JWTKEY=S3CUREK3Y
  LDAPBASE=dc=example,dc=org
  LDAPPORT=<ldap_port>
  LDAPHOST=<ldap_host>
  LDAPDN=cn=admin,dc=example,dc=org
  LDAPBINDPASSWORD=<ldapdn-password>
  LDAPSSL=true
  LDAPINSECURESKIP=true
  LDAPENABLE=true
  ADMINUSERNAME=admin
  ADMINPASSWORD=admin
  REGEXP=(.*nginx|.*-exporter|.*osxd)"
  JENKINSURL=<jenkins_url>
  ADMINS="<user1>,<user2>"
  JENKINSBUILDAPI="/job/..."
  JENKINSUSER=<jenkins_user>
  JENKINSPASS=<jenkins_password>
  S3BUCKET=<bucket>
  PRIVATEKEY=<location_of_key>
```
**Hint**

LDAP can be disable by turning `LDAPENABLE` to `false`

Jenkins can be disable by not setting any info regarding Jenkins

4. Run the solution
```
docker run -p 8080:8080 -v <private_key>:/<location_of_key>  --env-file <env_file> harbor-mgt-artifact 
```
5. Connect to `http://localhost:8080`