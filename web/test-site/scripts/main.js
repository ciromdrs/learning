var img = document.querySelector('img')

img.onclick = function (){
	var src = img.getAttribute('src')
	if(src === 'images/firefox.png'){
		img.setAttribute('src', 'images/firefox-2.jpg')
	} else {
		img.setAttribute('src', 'images/firefox.png')
	}
}


var btn = document.querySelector('button')
var h1 = document.querySelector('h1')

if (!localStorage.getItem('name')) {
	setUserName()
} else {
	var name = localStorage.getItem('name')
	h1.textContent = 'hello, '+name
} 


function setUserName(){
	var name = prompt("please enter your name.")
	localStorage.setItem('name',name)
	h1.textContent = 'hello, '+name

}

btn.onclick = function() {
	setUserName()
}