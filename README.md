Per creare i certificati necessari al protocollo TLS:

 openssl genrsa -out server.key 2048
 
 openssl ecparam -genkey -name secp384r1 -out server.key
 
 openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
