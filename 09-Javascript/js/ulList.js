function buildUL() {
	var div1 = document.getElementById("div1");
	var ul = document.createElement("ul");
	ul.innerHTML = "<li id=li1 onclick=display_alert('li1')>Sakura</li>" +
		"<li id=li2 onclick=display_alert('li2')>Ame</li>" +
		"<li id=li3 onclick=display_alert('li3')>Yuki</li>";
	div1.appendChild(ul);
}
