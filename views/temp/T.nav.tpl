{{define "nav"}}
  <body>
    <div class="container">
      <!--Logo area start-->
      <div style="float:left;width:80px">
        <img src="..." alt="..." class="img-thumbnail">
      </div>
      <!--logo end-->
      <!--corp's name start-->
      <div style="margin-left:150px">
            <h1>MyPacking Corporation Ltd.,Powered by Harry</h1>
      </div>
      <!--corp's name end-->
    </div>
    <!--nav area start-->
    <nav class="navbar navbar-default navbar-static-top">
      <div class="container" style="margin-left:50px">
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li {{if .IsHome}} class="active"{{end}}><a href="/">home</a></li>
            <li {{if .IsCategory}} class="active"{{end}}><a href="/category">Product</a></li>
            <li {{if .IsAbout}} class="active"{{end}}><a href="/introduce">About Us</a></li>
            <li {{if .IsNews}} class="active"{{end}}><a href="/news">News</a></li>
            <li {{if .IsContact}} class="active"{{end}}><a href="/contactus">Contact Us</a></li>
          </ul>
        </div>
      </div>
    </nav>
    <div class="bs-example" data-example-id="simple-jumbotron" >
        <div class="jumbotron" style="background-color:LightBlue">
          <h1>Hello, world!</h1>
          <p>This is a simple hero unit, a simple jumbotron-style component for calling extra attention to featured content or information.</p>
          <p></p>
        </div>
      </div>
  <!-- nav end-->

  {{end}}