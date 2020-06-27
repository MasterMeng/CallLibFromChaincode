```bash
peer chaincode install -n calc -v 1.0 -l golang -p github.com/chaincode/calc
peer chaincode instantiate -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n calc -l golang -v 1.0 -c '{"Args":["init"]}' -P 'OR ('\''Org1MSP.peer'\'','\''Org2MSP.peer'\'')'
peer chaincode query -C mychannel -n calc -c '{"Args":["add","1","2"]}'
```