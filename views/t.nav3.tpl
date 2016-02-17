{{define "nav3"}}
<nav class="navbar navbar-default" role="navigation">
        <!-- 导航头部 -->
        <div class="navbar-header">
          <!-- 移动设备上的导航切换按钮 -->
          <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse-example">
            <span class="sr-only">切换导航</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <!-- 品牌名称或logo -->
          <a class="navbar-brand" href="#">LOGO Here</a>
        </div>

        <!-- 导航项目 -->
        <div class="collapse navbar-collapse navbar-collapse-example">
          <!-- 一般导航项目 -->
          <ul class="nav navbar-nav">
            <li class="active"><a href="#">项目</a></li>
            <li><a href="#">需求</a></li>
            <!-- 导航中的下拉菜单 -->
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">管理 <b class="caret"></b></a>
              <ul class="dropdown-menu" role="menu">
                <li><a href="#">任务</a></li>
                <li><a href="#">待办</a></li>
                <li><a href="#">Bug</a></li>
                <li class="divider"></li>
                <li><a href="#">测试</a></li>
                <li><a href="#">用例</a></li>
              </ul>
            </li>
          </ul>
        </div><!-- END .navbar-collapse -->
      </nav>
{{end}}