const menuTray = document.getElementById("menuTray")

document.querySelector('#toggle').addEventListener("click", () => {
   if (menuTray.style.visibility === 'hidden') {
       menuTray.style.visibility = 'visible';
       document.getElementById("float-icon-arrow").classList.remove('fa-angle-double-up')
       document.getElementById("float-icon-arrow").classList.add('fa-angle-double-down')
   } else{
    menuTray.style.visibility = 'hidden';
    document.getElementById("float-icon-arrow").classList.remove('fa-angle-double-down')
       document.getElementById("float-icon-arrow").classList.add('fa-angle-double-up')
   }
})
