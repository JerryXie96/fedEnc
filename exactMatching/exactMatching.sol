pragma solidity = 0.5.4;
pragma experimental ABIEncoderV2;

contract ExactMatching {
    struct CipherText{
        bytes x;
        bytes y;
    }
    
    struct Index{
        CipherText CT;
        uint id;
    }
    
    Index[] indexCollection;
    
    uint[] result;
    
    // bool isPair=false;
    
    function store(CipherText memory ct, uint id) public {
        Index memory indexItem;
        indexItem.CT=ct;
        indexItem.id=id;
        indexCollection.push(indexItem);
    }
    
    function search(CipherText memory ct) public{
        bytes memory a=ct.x;
        bytes memory b=ct.y;
        for(uint i=0;i<indexCollection.length;i++){
            bytes memory cx=indexCollection[i].CT.x;
            bytes memory cy=indexCollection[i].CT.y;
            bytes memory input=abi.encodePacked(cx,b,a,cy);
            uint[1] memory output;
            uint length=input.length;
            assembly{
                if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
                    revert(0,0)
                }
            }
            if(output[0]!=0){
                // isPair=true;
                result.push(indexCollection[i].id);
            }
        }
    }
    
    // function storeOri(bytes memory x, bytes memory y, uint id) public {
    //     Index memory indexItem;
    //     indexItem.CT.x=x;
    //     indexItem.CT.y=y;
    //     indexItem.id=id;
    //     indexCollection.push(indexItem);
    // }
    
    // function searchOri(bytes memory a, bytes memory b) public{
    //     for(uint i=0;i<indexCollection.length;i++){
    //         bytes memory cx=indexCollection[i].CT.x;
    //         bytes memory cy=indexCollection[i].CT.y;

    //         bytes memory input=abi.encodePacked(cx,b,a,cy);
    //         uint[1] memory output;
    //         uint length=input.length;
    //         assembly{
    //             if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
    //                 revert(0,0)
    //             }
    //         }
    //         if(output[0]!=0){
    //             isPair=true;
    //             result.push(indexCollection[i].id);
    //         }
    //     }
    // }
    
    function getResult() public view returns(uint[] memory){
        return result;
    }
    
    // function getIndex() public view returns(Index[] memory){
    //     return indexCollection;
    // }
    
    // function getIsPair() public view returns(bool){
    //     return isPair;
    // }
    
    function clearRes() public {
        delete result;
    }
}