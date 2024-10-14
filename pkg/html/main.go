package html

var UI = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Telnet Client</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				margin: 0;
				padding: 0;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
			}
			.container {
				text-align: center;
				background-color: white;
				padding: 20px;
				border-radius: 10px;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				width: 300px;
			}
			h1 {
				margin-bottom: 20px;
			}
			label {
				display: block;
				margin-bottom: 10px;
				font-weight: bold;
			}
			input[type="text"] {
				width: 100%;
				padding: 10px;
				margin-bottom: 20px;
				border: 1px solid #ccc;
				border-radius: 5px;
				box-sizing: border-box;
			}
			input[type="submit"] {
				padding: 10px 20px;
				background-color: #4CAF50;
				color: white;
				border: none;
				border-radius: 5px;
				cursor: pointer;
			}
			input[type="submit"]:hover {
				background-color: #45a049;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Telnet Connection</h1>
			<form action="/telnet" method="POST">
				<label for="ip">IP Address:</label>
				<input type="text" id="ip" name="ip" placeholder="Enter IP" required>

				<label for="port">Port:</label>
				<input type="text" id="port" name="port" placeholder="Enter Port" required>

				<input type="submit" value="Connect">
			</form>
		</div>
	</body>
	</html>
	`
