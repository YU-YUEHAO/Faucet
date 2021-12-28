
if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!')
      getAccount();
      ethereum.on('accountsChanged', function (accounts) {
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
      gettodatdate();
      getalldate();
  }
function getbalance(){
   $.ajax({
       method:"post",
       url:"http://localhost:1111/ygh/cAmount",
       data:{account:account},
       success:function (data){
           // console.log(data.data)
           document.getElementById("ercNumber").innerHTML=data.data+"ygh";

       },
       err:function (err){
           console.log(err)
       }
   })
}
var t;
function gettodaytoken(){
      $.ajax({
          method:"post",
          url: "http://localhost:1111/ygh/gettoday",
          data: {account,account},
          success:function (data){
              t=data.data;
              document.getElementById("graph7").innerHTML=data.data+"ygh"
          },
          err:function (err){
              console.log(err)
          }
      })
}
function getuserall(){
    $.ajax({
        method:"post",
        url:"http://localhost:1111/ygh/useralltoke",
        data:{account:account},
        success:function (data){
            // console.log(data.data);
            document.getElementById("graph8").innerHTML=data.data+"ygh"

        },
        err:function (data){
            console.log(data)
        }
    })
}


function dotransfer(){
    gettodaytoken();
    if (t>=3){
        alert("today is to many requst,welcome tomorrow")

    }else{
    $.ajax({
        method:"post",
        url:"http://localhost:1111/ygh/dotransfer",
        data:{account:account},
        success:function (data){
            console.log(data);
            if (data.code==0){
            console.log("tranfer1"+data)
            gettodaytoken();
            getuserall();
            getnewtime();
            getuserbalance();
            gettodatdate();
            getalldate();
            alert("本次获取成功")
                console.log("okoko");
            window.location.href="http://localhost:1111";
            }else if (data.code==1){
                console.log("tranfer2"+data)
                 alert(data.message)
            }
        },
        err:function (data){
            console.log(data);
            alert("交易有错")
        }
    })
    }
}
function  getnewtime(){
       $.ajax({
            method:"post",
            url:"http://localhost:1111/ygh/gettime",
            data:{account,account},
           success:function (data){
                // console.log(data.data)
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
             // console.log(data.data);
             document.getElementById("graph9").innerHTML=data.data+"ygh";

          },
          err:function (data){
                console.log(data)
          }
      })
}

function gettodatdate(){
      $.ajax({
          method:"post",
          url:"http://localhost:1111/ygh/todaydata",
          data:{account:account},
          success:function (data){
              // console.log(data.data);
              document.getElementById("today").innerHTML=data.data+"ygh"
          },
          err:function (data){
              console.log(data)
          }
      })
}
function getalldate(){
      $.ajax({
          method:"post",
          url:"http://localhost:1111/ygh/alldata",
          data:{account:account},
          success:function (data){
              // console.log(data.data);
              document.getElementById("alldate").innerHTML=data.data+"ygh"
          },
          err:function (data){
              console.log(data);
          }
      })
}

