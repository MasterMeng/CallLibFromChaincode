From hyperledger/fabric-ccenv:1.4.6
COPY payload/calc.h /usr/local/include
#COPy payload/libcalc.a /usr/local/lib 
COPy payload/libcalc.so /usr/local/lib 
