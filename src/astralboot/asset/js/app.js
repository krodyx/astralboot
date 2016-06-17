source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span>Error : </span>'+value+'</div>';
    return alertBox;
}

source.addEventListener('info',function(e){
//    console.log(e);
    document.getElementById("flag").classList.toggle("notif");
})

source.addEventListener('alert',function(e){
    var d1 = document.getElementById("alert");
    d1.insertAdjacentHTML('beforeend',AlertBox(e.data));
})
