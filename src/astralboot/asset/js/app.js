source = new EventSource('/events');
source.addEventListener('info',function(e){
//    console.log(e);
    document.getElementById("flag").classList.toggle("notif")
})

