// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract faucet{
    // address admin=0x00DC6E8B60fA02a5d83e525BBEf3240E8ea54dc5;

    struct users{
         address _address;
         uint  transfercount;//请求的账户
         uint  transfertime;//最近一次交易时间
         uint  tokenaccount;//一共获取的币数
         uint todaytokenaccount;//当天获取的币数
    }

    mapping (address=>users) public usertable;
 
    //用户可以通过合约转账1个ether
    function dotransfer(address add)  public{
         require(address(this).balance>=0.01 ether,"no money");
         uint currentTime =block.timestamp;
         if(usertable[add].tokenaccount >=100)
         revert("you have too many token");
         if(usertable[add].todaytokenaccount >=5)
         revert("today get token limit 5");
         if(currentTime-usertable[add].transfertime>=1 days)
         refreshtodaytoken(add);
         if(currentTime-usertable[add].transfertime<=1 days && usertable[add].transfercount>=3)
         revert("today is to many requst,welcome tomorrow");
         if (currentTime - usertable[add].transfertime <=10 seconds)
         revert("requesting too often ");
         uint u = 1 ether;
         payable(add).transfer(u);
         usertable[add].transfercount++;
         usertable[add].tokenaccount++;
         usertable[add].todaytokenaccount++;
         usertable[add].transfertime=block.timestamp;
  }

  //刷新当天的货币数量
  function refreshtodaytoken(address add) public payable {
              usertable[add].todaytokenaccount=0;
  }
   //获取当天得到的货币数量
  function gettodaytoken(address add)public view returns(uint){
                 return usertable[add].todaytokenaccount;
  }

//获取账户得币数量
 function gettransfercount(address add) public view returns(uint){
         return usertable[add].transfercount;
 }
 //获取账户最新交易时间
 function gettransfertime(address add)public view returns(uint){
     return  usertable[add].transfertime;
 }
 //获取合约余额 
 function getBalancecontract()public view returns(uint){
     return address(this).balance/10**18;
}
//获取本用户余额
 function getBalanceUser()public view returns(uint){
         return msg.sender.balance/10**18;
    
 }
  fallback() external payable {}
  receive() external payable {}
}