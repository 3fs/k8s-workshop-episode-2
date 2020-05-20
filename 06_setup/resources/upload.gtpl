<html>

<head>
	<title>Upload file</title>
	<style>
		body {
			font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
			display: flex;
			height: 100%;
			margin: 0;
			justify-content: center;
			align-items: center;
		}

		.inputfile {
			width: 0.1px;
			height: 0.1px;
			opacity: 0;
			overflow: hidden;
			position: absolute;
			z-index: -1;
		}

		.inputfile+label {
			max-width: 80%;
			font-size: 1.25rem;
			font-weight: 700;
			text-overflow: ellipsis;
			white-space: nowrap;
			cursor: pointer;
			display: inline-block;
			overflow: hidden;
			padding: 0.625rem 1.25rem;
		}

		.inputfile:focus+label,
		.inputfile.has-focus+label {
			outline: 1px dotted #000;
			outline: -webkit-focus-ring-color auto 5px;
		}

		.inputfile+label svg {
			width: 1em;
			height: 1em;
			vertical-align: middle;
			fill: currentColor;
			margin-top: -0.25em;
			margin-right: 0.25em;
		}

		.inputfile+label {
			color: #172032;
		}

		.inputfile:focus+label,
		.inputfile.has-focus+label,
		.inputfile+label:hover {
			color: #172032;
		}

		.inputfile+label figure {
			width: 3rem;
			height: 3rem;
			border-radius: 50%;
			background-color: #ffd400;
			display: block;
			padding: 2rem;
			margin: 0 auto 1rem;
		}

		.inputfile:focus+label figure,
		.inputfile.has-focus+label figure,
		.inputfile+label:hover figure {
			background-color: #f4cb00;
		}

		.inputfile+label svg {
			width: 100%;
			height: 100%;
			fill: #000;
		}

		#fileList {
			list-style: none;
			margin: 0;
			padding: 0;
		}

		#fileList:empty+input {
			display: none;
		}

		input[type="submit"] {
			background-color: #172032;
			border: none;
			border-radius: 999rem;
			color: white;
			padding: 1rem 2rem;
			font-size: 1rem;
			text-decoration: none;
			margin: 1rem;
			cursor: pointer;
		}

		input[type="submit"]:focus,
		input[type="submit"]:hover {
			opacity: .8;
		}

		form {
			background: #f5f5f5;
			padding: 3rem;
			text-align: center;
			max-width: 400px;
			border-radius: .3rem;
		}
	</style>
</head>

<body>
	<form enctype="multipart/form-data" action="/upload" method="post">
		<input type="file" name="uploadfile" id="uploadfile" class="inputfile" multiple onchange="updateList()" />
		<label for="uploadfile">
			<figure><svg xmlns="http://www.w3.org/2000/svg" width="20" height="17" viewBox="0 0 20 17">
					<path
						d="M10 0l-5.2 4.9h3.3v5.1h3.8v-5.1h3.3l-5.2-4.9zm9.3 11.5l-3.2-2.1h-2l3.4 2.6h-3.5c-.1 0-.2.1-.2.1l-.8 2.3h-6l-.8-2.2c-.1-.1-.1-.2-.2-.2h-3.6l3.4-2.6h-2l-3.2 2.1c-.4.3-.7 1-.6 1.5l.6 3.1c.1.5.7.9 1.2.9h16.3c.6 0 1.1-.4 1.3-.9l.6-3.1c.1-.5-.2-1.2-.7-1.5z" />
				</svg></figure> <span>Choose a file&hellip;</span>
		</label>
		<div id="fileList"></div>
		<input type="submit" value="Upload" />
	</form>
	<script>
		function updateList() {
			var input = document.querySelector('input[name="uploadfile"]');
			var output = document.getElementById('fileList');
			output.innerHTML = '<ul>';
			for (var i = 0; i < input.files.length; ++i) {
				output.innerHTML += '<li>' + input.files.item(i).name + '</li>';
			}
			output.innerHTML += '</ul>';
		}
	</script>
</body>

</html>
