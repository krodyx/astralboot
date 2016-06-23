source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

function Acknowledge(value){
    var alertBox = '<div id="'+value.UUID+'" class="alert-box error"><span></span>'+value.Status+' <a class="pure-button" onclick="FetchClick(\''+value.UUID+'\');" >ACK</a></div>';
    return alertBox;
}

function FetchClick(id){
    console.log(id);
    fetch('ack/'+id,{method: 'get' }).then(function(response){
        console.log(response);
    });
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
    var data = JSON.parse(e.data);
    d1.insertAdjacentHTML('beforeend',AlertBox(data));
});

source.addEventListener('ack',function(e){
    console.log(e);
    var d1 = document.getElementById("alert");
    var data = JSON.parse(e.data);
    d1.insertAdjacentHTML('beforeend',Acknowledge(data));
});
//source.onerror = function(e){
//   console.log(e);
//  var d1 = document.getElementById("alert");
// d1.insertAdjacentHTML('beforeend',AlertBox('fail'));
//    source.close();
//};
