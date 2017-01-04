# Nmea tracker

## Protocol
As it's first version, protocol is extremely simple.
Clients should connect by tcp to nmea.scukonick.rocks:3999 and send message/messages.

The first line should be like HTTP header, containing <b>token</b> of device sending this data.
Using this token you can log in on the site and see last history of the device.<br>
The next lines are just NMEA lines received from GPS.

Lines could be separated by `<CR><LF>` as well as by `<LF>`.

There is no 'end-of-conversation' symbol, it's ok just to close connection from the device.
If device is not closing connection the server would close it after some timeout.

Example conversation:

```
Token: 12312412412
$GPRMC,002454,A,3553.5295,N,13938.6570,E,0.0,43.1,180700,7.1,W,A*3F
$GPRMB,A,,,,,,,,,,,,A,A*0B
$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*7F
$GPGSA,A,3,01,04,07,16,20,,,,,,,,3.6,2.2,2.7*35
$GPGSV,3,1,09,01,38,103,37,02,23,215,00,04,38,297,37,05,00,328,00*70
$GPGSV,3,2,09,07,77,299,47,11,07,087,00,16,74,041,47,20,38,044,43*73
$GPGSV,3,3,09,24,12,282,00*4D
$GPGLL,3553.5295,N,13938.6570,E,002454,A,A*4F
$GPBOD,,T,,M,,*47
$PGRME,8.6,M,9.6,M,12.9,M*15
$PGRMZ,51,f*30
$HCHDG,101.1,,,7.1,W*3C
$GPRTE,1,1,c,*37
$GPRMC,002456,A,3553.5295,N,13938.6570,E,0.0,43.1,180700,7.1,W,A*3D
$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*7F
$GPGSA,A,3,01,04,07,16,20,,,,,,,,3.6,2.2,2.7*35
$GPGSV,3,1,09,01,38,103,37,02,23,215,00,04,38,297,37,05,00,328,00*70
$GPGSV,3,2,09,07,77,299,47,11,07,087,00,16,74,041,47,20,38,044,43*73
$GPGSV,3,3,09,24,12,282,00*4D
$GPGLL,3553.5295,N,13938.6570,E,002454,A,A*4F
$HCHDG,101.1,,,7.1,W*3C
$GPRTE,1,1,c,*37
$GPRMC,002456,A,3553.5295,N,13938.6570,E,0.0,43.1,180700,7.1,W,A*3D
```
