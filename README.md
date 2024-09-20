# OVERVIEW 
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/uniborg.svg)](https://pkg.go.dev/paepcke.de/uniborg) 
[![Go Report Card](https://goreportcard.com/badge/paepcke.de/uniborg)](https://goreportcard.com/report/paepcke.de/uniborg) 

[paepcke.de/uniborg](https://paepcke.de/uniborg/)

# UNIBORG 

- secure backup & orchestration for [ui.com](https://ui.com/) controller appliances
  
# EXAMPLE 
```
UNI_TARGETS="uni001.lan,unni002.lan,uni003.lan" UNI_APIKEY="..." UNI_APISECRET="..." go run paepcke.de/uniborg/cmd/uniborg@latest
```

# SUPPORTED OPTIONS 

```
# REQUIRED: 
- UNI_TARGETS   - list of Unifi Controller Target Server to Backup (comma separated)
- UNI_APIKEY    - Unifi Controller Backup User APIKEY
- UNI_APISECRET - Unifi Controller User APISECRET

# OPTIONAL:
- UNI_TLSKEYPIN - Unifi TLS MitM proof Certificate Keypin [string]
- UNI_DAEMON    - run app in daemon mode, never quit, fetch once every hour [bool: defaults to 'false']
- UNI_NOGIT     - do not create & update local git version repo [bool: defaults to 'false']
- UNI_NOSSL     - do not verify SSL Certificates [bool: defaults to *'true'*, SSL SystemCertStore is pointless, use UNI_TLSKEYPIN!]

```
# OPTIONS FAQ

```
- Enviroment Variables bools must set exactly case senitive to 'true' to enable, anything else will default to false.
- Clear text HTTP protocol is not supported, switch on HTTPS for your admin interface (self-signed certificates will do)
- ATT: HTTPS chain verification via system os trust store(s) clusterfuck is disabled by default (UNI_NOSSL='true'), use UNI_TLSKEYPIN !
```

# HOW TO INSTALL

```
go install paepcke.de/uniborg/cmd/uniborg@latest
```

### DOWNLOAD (prebuild)

[github.com/paepckehh/uniborg/releases](https://github.com/paepckehh/uniborg/releases)


# STATUS

 - -=# WORK IN PROGRESS - DO NOT USE THIS REPO - MAY IMPLODE ANYTIME #=- 
 - -=# WORK IN PROGRESS - DO NOT USE THIS REPO - MAY IMPLODE ANYTIME #=- 
 - -=# WORK IN PROGRESS - DO NOT USE THIS REPO - MAY IMPLODE ANYTIME #=- 

# TIMELINE 

 - 2024
    - September -> Internal Use Only (codesamples)
    - October   -> Offical Testphase 
    - November  -> Offical Pilot Phase
    - December  -> Release Candidate
 - 2025
    - January - > First Public Release & Native [NixOS](https://github.com/nixos) support

# CONTRIBUTION

Yes, Please! PRs Welcome! 

# SPONSORS 



