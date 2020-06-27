# 在Chaincode中调用C库  

<font color="red">本作品采用[署名-非商业性使用-相同方式共享 4.0 国际 (CC BY-NC-SA 4.0)](https://creativecommons.org/licenses/by-nc-sa/4.0/deed.zh)进行许可，使用时请注明出处。</font>

在fabric的chaincode开发时，有时候需要用到第三方库提供的功能。这些库有些是没有go的实现，或开发者只提供了库，这时候就需要从chaincode中调用第三方库。而fabric的chaincode都是运行在docker容器中，此时就需要将需要用的第三方库编译到dockers容器中。  

fabric官方支持四种语言的chaincode开发，分别为[go](https://github.com/hyperledger/fabric-chaincode-go)、[java](https://github.com/hyperledger/fabric-chaincode-java)、[node](https://github.com/hyperledger/fabric-chaincode-node)和[evm](https://github.com/hyperledger/fabric-chaincode-evm)。这里以**Go**为例，介绍如何在chaincode中调用第三方库。示例中用到的所有文件都可以从[这里](https://github.com/MasterMeng/CallLibFromChaincode)找到。  

本文使用[fabric-samples v1.4.6](https://github.com/hyperledger/fabric-samples)中提供的**first-network**来搭建fabric网络。关于fabric网络开发环境的搭建，详见我的另一篇[文章](https://www.cnblogs.com/lianshuiwuyi/p/11819131.html)，这里就不再赘述了。  

## 1 使用静态库  

在fabric网络中，默认启用静态库支持。所以使用静态库时，只需将所需的静态库编译到fabric镜像中即可。研究过fabric源码你会发现，fabric网络的基本配置都是通过**core.yaml**文件提供的，chaincode的配置也不例外。  

### 1.1 重编镜像

在默认情况下，chaincode会在**fabric-ccenv:latest**中编译成可执行程序，所以我们需要重写编译fabric-ccenv镜像，将chaincode需要的所有依赖都打包到镜像中，并将新生成的镜像tag为`latest`。  

Dockerfile的内容如下：  

```
From hyperledger/fabric-ccenv:1.4.6
COPY payload/calc.h /usr/local/include
COPy payload/libcalc.a /usr/local/lib
```   

### 1.2 修改配置

在ccenv中生成的可执行程序会与`chaincode.LANG.runtime`指定的镜像一起生成chaincode运行的容器，例如，golang的chaincode使用的时*fabric-baseos:\$(ARCH)-\$(BASE_VERSION)*镜像。这里我把替换成了*fabric-ccenv:latest*，这样可以保证chaincode运行时的所有依赖都能正常加载。  

### 1.3 启动网络

本示例中使用的库只提供了简单的加减乘除功能。拿到后只需放到`fabric-samples/chaincode/`路径下，将修改后的`core.yaml`映射到所有peer节点的`/etc/hyperledger/fabric`路径下，即可使用`first-network`下的`byfn.sh`脚本启动网络。chaincode使用方法如下：  

``` bash
peer chaincode install -n calc -v 1.0 -l golang -p github.com/chaincode/calc
peer chaincode instantiate -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n calc -l golang -v 1.0 -c '{"Args":["init"]}' -P 'OR ('\''Org1MSP.peer'\'','\''Org2MSP.peer'\'')'
peer chaincode query -C mychannel -n calc -c '{"Args":["add","1","2"]}'
```  

## 2 使用动态库  

### 2.1 重编镜像  

与使用静态库相同，使用动态库时首先还是需要从新编译ccenv镜像，将所有的依赖都打包到镜像中。  

### 2.2 修改配置  

将`chaincode.golang.runtime`指定的镜像替换成`fabric-ccenv:latest`。因为使用的go的chaincode，还需要将`chaincode.golang.dynamicLink`改为`true`来启用动态库链接。  

### 2.3 启动网络  

启动网络的操作与静态库的操作一样。