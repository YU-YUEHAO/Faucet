
if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!')
      getAccount();
      ethereum.on('accountsChanged', function (accounts) {
          // Time to reload your interface with accounts[0]!
          console.log("现在的账户是:", accounts[0]);
          account=accounts[0];
          $("#address").html(accounts[0]);
      });
  } else {
    alert('plase install the MetaMask')
  }
  var account;
  async function getAccount() {
      yghaccounts = await ethereum.request({ method: 'eth_requestAccounts' });
      account = yghaccounts[0];
     $("#address").html(account);
     console.log(account)
    getbalance();
     gettodaytoken();
     getuserall();
      getnewtime();
      getuserbalance();
  }
function getbalance(){
   $.ajax({
       method:"post",
       url:"http://localhost:1111/ygh/cAmount",
       data:{account:account},
       success:function (data){
           console.log(data.data)
           document.getElementById("ercNumber").innerHTML=data.data+"ygh";

       },
       err:function (err){
           console.log(err)
       }
   })
}
function gettodaytoken(){
      $.ajax({
          method:"post",
          url: "http://localhost:1111/ygh/gettoday",
          data: {account,account},
          success:function (data){
              console.log(data)
              document.getElementById("graph7").innerHTML=data.data+"ygh"
          },
          err:function (data){
              console.log(data)

          }
      })
}
function getuserall(){
    $.ajax({
        method:"post",
        url:"http://localhost:1111/ygh/useralltoke",
        data:{account:account},
        success:function (data){
            console.log(data.data);
            document.getElementById("graph8").innerHTML=data.data+"ygh"

        },
        err:function (data){
            console.log(data)
        }
    })
}

function dotransfer(){
    $.ajax({
        method:"post",
        url:"http://localhost:1111/ygh/dotransfer",
        data:{account:account},
        success:function (data){
            console.log(data)
            gettodaytoken();
            getuserall();
            getnewtime();
            getuserbalance();
            alert("本次获取成功")
            window.location.href="http://localhost:1111";
        },
        err:function (data){
            console.log(data);
            alert("交易有错")
        }
    })
}
function  getnewtime(){
       $.ajax({
            method:"post",
            url:"http://localhost:1111/ygh/gettime",
            data:{account,account},
           success:function (data){
                console.log(data.data)
               console.log(typeof(data.data))
               if (data.data=="1970-01-01T08:00:00+08:00"){
                   document.getElementById("tranfertime").innerHTML="无最近交易时间"
               }else {

                document.getElementById("tranfertime").innerHTML=data.data;
               }
           }
       })
}
function getuserbalance(){
      $.ajax({
          method:"post",
          url:"http://localhost:1111/ygh/getbalaceUser",
          data:{account:account},
          success:function (data){
             console.log(data.data);
             document.getElementById("graph9").innerHTML=data.data+"ygh";

          },
          err:function (data){
                console.log(data)
          }
      })
}


