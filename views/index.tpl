<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  

</head>

<body>
  <div class="container">
      <h1>上传文件</h1>
      <form method="post" action="/upload" enctype="multipart/form-data">
        <div class="form-group">
          <!-- <label>文章附件：</label> -->
          <input type="file" class="form-control" name="aaaa">
        </div>
        <button type="submit" class="btn btn-default">上传</button>
      </form>
    </div>

    <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</body>
</html>
