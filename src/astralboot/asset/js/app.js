source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

function Acknowledge(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'<a class="pure-button" href="#">ACK</a></div>';
    return alertBox;
}

source.addEventListener('info',function(e){
//    console.log(e);
    document.getElementById("flag").classList.toggle("notif");
});

source.addEventListener('tick',function(e){
    //console.log(e);
});

source.addEventListener('status',function(e){
    var data = JSON.parse(e.data);
    console.log(data);
    document.getElementById(data.Name).classList.toggle(data.Status);
});

source.addEventListener('alert',function(e){
    var d1 = document.getElementById("alert");
    d1.insertAdjacentHTML('beforeend',AlertBox(e.data));
});

source.addEventListener('ack',function(e){
    console.log(e);
    var d1 = document.getElementById("alert");
    d1.insertAdjacentHTML('beforeend',Acknowledge(e.data));
});
//source.onerror = function(e){
//   console.log(e);
//  var d1 = document.getElementById("alert");
// d1.insertAdjacentHTML('beforeend',AlertBox('fail'));
//    source.close();
//};
