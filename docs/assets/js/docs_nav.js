const menuTray = document.getElementById("menuTray")


document.querySelector('#toggle').addEventListener("click", () => {
   if (menuTray.style.visibility === 'hidden') {
       menuTray.style.visibility = 'visible';
       document.getElementById("float-icon-arrow").classList.remove('fa-angle-double-right')
       document.getElementById("float-icon-arrow").classList.add('fa-angle-double-left')
   }else{
    menuTray.style.visibility = 'hidden';
    document.getElementById("float-icon-arrow").classList.remove('fa-angle-double-left')
       document.getElementById("float-icon-arrow").classList.add('fa-angle-double-right')
   }
})

$(window).resize(function() { 
    if (menuTray.style.visibility === 'visible' && window.innerWidth > 575) {
        menuTray.style.visibility = 'hidden'
        document.getElementById("float-icon-arrow").classList.remove('fa-angle-double-left')
       document.getElementById("float-icon-arrow").classList.add('fa-angle-double-right')
    }    });