source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span>Error : </span>'+value+'</div>';
    return alertBox;
}

source.addEventListener('info',function(e){
//    console.log(e);
    document.getElementById("flag").classList.toggle("notif");
});

source.addEventListener('tick',function(e){
    console.log(e);
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

//source.onerror = function(e){
//   console.log(e);
//  var d1 = document.getElementById("alert");
// d1.insertAdjacentHTML('beforeend',AlertBox('fail'));
//    source.close();
//};
