package main

// http request needs a request line, body,  method, header

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	listner , err := net.Listen("tcp", ":8011")
	if err !=  nil {
		log.Fatalln(err.Error())
	}
	defer listner.Close()

	for {
		connection , err := listner.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		// spin up go ruutine to handle the request
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn){
	defer connection.Close()
	// read request
	request(connection)
}

func request(connection net.Conn){
	i := 0 // first time through will yeild us request method
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if i == 0 {
			multiPlexer(connection, line)
		}
		if line == "" {
			break
		}
		i++
	}
}

func multiPlexer(connection net.Conn, line string){
				// request line
				method := strings.Fields(line)[0]
				url := strings.Fields(line)[1]
				fmt.Println("****METHOD", method)
				fmt.Println("****URL", url)

				if method == "GET" && url == "/" {
					index(connection)
				}
				if method == "GET" && url =="/about" {
					about(connection)
				}

}

func index(connection net.Conn){
	body := `
	<!-- DOCTYPE html--->
	<html lang="en">
	<body>
	<canvas id="canvas" style="height:100%; width:100%;"></canvas>
	</body>
	<script  type="text/javascript">
	console.log("here")


	window.onload = function(){
		console.log("here 1")
		let c = document.getElementById('canvas');
	let ctx = c.getContext('2d');
	let i = 0;
	
	ctx.canvas.width = window.innerWidth;
	ctx.canvas.height = window.innerHeight;
	ctx.canvas.style.background = '#'+Math.floor(Math.random()*16777215).toString(16);
	ctx.canvas.style.position = 'fixed';
	
	let min = (canvas.width < canvas.height) ? canvas.width : canvas.height;
	
	function drawCircle() {
	  i++;
	  let weight = Math.pow(Math.random(), 2);
	  let seed1 = Math.floor(Math.floor(weight * ((min / 2) - 0 + 1)));
	  let seed2 = Math.floor(Math.random() * canvas.height);
	  let seed3 = Math.floor(Math.random() * canvas.width);
	  let primary = '#'+Math.floor(Math.random()*16777215).toString(16);
	  let secondary = '#'+Math.floor(Math.random()*16777215).toString(16);
	  let arc = {
		x: seed3 - (seed1 / 2),
		y: seed2 - (seed1 / 2),
		r: seed1,
		shadow: !(Math.random()+0.5|0),
		blur: Math.random() * (10 + 1),
		color: (!(Math.random()+0.5|0)) ? primary : secondary,
		stroke: Math.random() * (15 + 1),
		fill: !(Math.random()+0.5|0)
	  };
	  ctx.beginPath();
	  ctx.arc(arc.x, arc.y, arc.r, Math.PI * 0, Math.PI * 2);
	  
	  if (arc.fill) {
		ctx.fillStyle = arc.color;
		ctx.fill();
	  } else {
		ctx.strokeStyle = arc.color;
		ctx.lineWidth = arc.stroke;
		ctx.stroke();
	  }
	  if (i < 2000) {
		setTimeout(function(){
		  drawCircle();
		}, 10);
	  }
	}
	
	drawCircle();
	}
	</script>
	</html>`
	fmt.Fprint(connection, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(connection, "Content-Type: text/html \r\n")	
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection,	body)	
}

func about(connection net.Conn){
	body := `
	<!-- DOCTYPE html--->
	<html lang="en">
	<body class="main">
	<canvas id="canvas" style="height:100%; width:100%;"></canvas>
	</body>
	<script  type="text/javascript">
	console.log("here")


	window.onload = function(){
		console.log("here 1")
		const b = document.getElementsByClassName('main')[0]
		b.style.background = "white"
		let c = document.getElementById('canvas');
	let ctx = c.getContext('2d');
	let i = 0;
	
	ctx.canvas.width = window.innerWidth;
	ctx.canvas.height = window.innerHeight;
	ctx.canvas.style.background = '#'+Math.floor(Math.random()*16777215).toString(16);
	ctx.canvas.style.position = 'fixed';
	
	let min = (canvas.width < canvas.height) ? canvas.width : canvas.height;
	
	function drawCircle() {
	  i++;
	  let weight = Math.pow(Math.random(), 2);
	  let seed1 = Math.floor(Math.floor(weight * ((min / 2) - 0 + 1)));
	  let seed2 = Math.floor(Math.random() * canvas.height);
	  let seed3 = Math.floor(Math.random() * canvas.width);
	  let primary = '#'+Math.floor(Math.random()*16777215).toString(16);
	  let secondary = '#'+Math.floor(Math.random()*16777215).toString(16);
	  let arc = {
		x: seed3 - (seed1 / 2),
		y: seed2 - (seed1 / 2),
		r: seed1,
		shadow: !(Math.random()+0.5|0),
		blur: Math.random() * (10 + 1),
		color: (!(Math.random()+0.5|0)) ? primary : secondary,
		stroke: Math.random() * (15 + 1),
		fill: !(Math.random()+0.5|0)
	  };
	  ctx.beginPath();
	  ctx.arc(arc.x, arc.y, arc.r, Math.PI * 0, Math.PI * 2);
	  
	  if (arc.fill) {
		ctx.fillStyle = "#00000";
		ctx.fill();
	  } else {
		ctx.strokeStyle = arc.color;
		ctx.lineWidth = arc.stroke;
		ctx.stroke();
	  }
	  if (i < 2000) {
		setTimeout(function(){
		  drawCircle();
		}, 10);
	  }
	}
	
	drawCircle();
	}
	</script>
	</html>`
	fmt.Fprint(connection, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(connection, "Content-Type: text/html \r\n")	
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection,	body)	
}
