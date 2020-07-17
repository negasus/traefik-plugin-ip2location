# ip2location

Plugin for getting information from ip2location database and pass it to request headers

## Configuration

To configure this plugin you should add its configuration to the Traefik dynamic configuration as explained [here](https://docs.traefik.io/getting-started/configuration-overview/#the-dynamic-configuration).
The following snippet shows how to configure this plugin with the File provider in TOML and YAML: 

Static:

```yaml
experimental:
  pilot:
    token: xxx

  plugins:
    ip2location:
      modulename: github.com/negasus/traefik-plugin-ip2location
      version: v0.1.0
```

Dynamic:

```yaml
http:
  middlewares:
   my-plugin:
      plugin:
        ip2location:
          filename: /path/to/database.bin
          fromHeader: X-User-IP # optional
          headers:
            CountryShort: X-GEO-CountryShort
            CountryLong: X-GEO-CountryLong
            Region: X-GEO-Region
            City: X-GEO-City
            Isp: X-GEO-Isp
            Latitude: X-GEO-Latitude
            Longitude: X-GEO-Longitude
            Domain: X-GEO-Domain
            Zipcode: X-GEO-Zipcode
            Timezone: X-GEO-Timezone
            Netspeed: X-GEO-Netspeed
            Iddcode: X-GEO-Iddcode
            Areacode: X-GEO-Areacode
            Weatherstationcode: X-GEO-Weatherstationcode
            Weatherstationname: X-GEO-Weatherstationname
            Mcc: X-GEO-Mcc
            Mnc: X-GEO-Mnc
            Mobilebrand: X-GEO-Mobilebrand
            Elevation: X-GEO-Elevation
            Usagetype: X-GEO-Usagetype
```

### Options

#### Filename (`filename`)

*Required*

The path to ip2location database file (in binary format)

#### FromHeader (`fromHeader`)

*Default: empty*

If defined, IP address will be obtained from this HTTP header

#### Headers (`headers`)

*Default: empty*

Define the HTTP Header name if you want to pass any of the parameters