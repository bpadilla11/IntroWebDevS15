var quotes = [
	"People are satisfied to judge things by their own narrow experience, never knowing of the wide world outside.",
	"Perseverance is better than defeat.",
	"Nothing ventured, nothing gained."
]

var count = 1

function transition() {
	if(count == 1) {
		document.querySelector("#bg1").className = "bg-active";
		document.querySelector("#bg2").className = "";
		document.querySelector("#bg3").className = "";
		count++;
	}
	else if(count == 2) {
		document.querySelector("#bg1").className = "";
		document.querySelector("#bg2").className = "bg-active";
		document.querySelector("#bg3").className = "";
		count++;
	}
	else {
		document.querySelector("#bg1").className = "";
		document.querySelector("#bg2").className = "";
		document.querySelector("#bg3").className = "bg-active";
		count = 1;
	}
	var i = Math.floor((Math.random() * 3));
	document.querySelector("#quote").innerHTML = quotes[i];
}

setInterval(transition, 4000);
