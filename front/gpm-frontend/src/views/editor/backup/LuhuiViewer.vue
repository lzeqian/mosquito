<style>
    .literally {
        width: 100%;
        height: 100%;
        position: relative;
    }

    #main_menu ul li a .rightarrowclass {
        background: url(https://luhuiphotoshop.netlify.app/img/right.gif) no-repeat;
        width: 8px;
        height: 11px;
    }

    #main_menu ul li a .rightarrowclass {
        background: url(https://luhuiphotoshop.netlify.app/img/right.gif) no-repeat;
        width: 8px;
        height: 11px;
    }
</style>
<template>

    <div ref="element" id="my-drawing" style="height: 100%;overflow: auto;margin-left: 5px;background-color: grey;margin-top:25px">

        <div id="drop_zone" >
            <div style="margin-top:70px;">
                <span style="font-size: 40px;">DROP</span>
                or <input type="file" onchange="upload_browse(this.files)">
            </div>
        </div>
        <div id="canvas_container">
            <div id="menu_left">
                <div id="menu_left_container"><a title="Select object tool"
                                                 style="background: #989898 url(&#39;img/all.png&#39;) no-repeat 7px 2px;"
                                                 class="" onclick="return TOOLS.action(&#39;select_tool&#39;);"
                                                 id="select_tool" href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Select area tool"
                       style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -46px 5px;"
                       onclick="return TOOLS.action(&#39;select_square&#39;);" id="select_square"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Magic Wand Tool"
                       style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -149px -48px;"
                       onclick="return TOOLS.action(&#39;magic_wand&#39;);" id="magic_wand"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Erase" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -97px 4px;"
                       onclick="return TOOLS.action(&#39;erase&#39;);" id="erase"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Fill" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -147px 3px;"
                       onclick="return TOOLS.action(&#39;fill&#39;);" id="fill"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Pick Color" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -197px 3px;"
                       onclick="return TOOLS.action(&#39;pick_color&#39;);" id="pick_color"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Pencil" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -247px 3px;"
                       onclick="return TOOLS.action(&#39;pencil&#39;);" id="pencil"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Draw line" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -297px 3px;"
                       onclick="return TOOLS.action(&#39;line&#39;);" id="line"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Draw letters" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -347px 4px;"
                       onclick="return TOOLS.action(&#39;letters&#39;);" id="letters"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Draw rectangle"
                       style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -397px 5px;"
                       onclick="return TOOLS.action(&#39;draw_square&#39;);" id="draw_square"
                       href="https://luhuiphotoshop.netlify.app/#" class="active"></a>
                    <a title="Draw circle" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -447px 5px;"
                       onclick="return TOOLS.action(&#39;draw_circle&#39;);" id="draw_circle"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Brush" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -494px 3px;"
                       onclick="return TOOLS.action(&#39;brush&#39;);" id="brush"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Blur tool" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -245px -48px;"
                       onclick="return TOOLS.action(&#39;blur_tool&#39;);" id="blur_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Sharpen tool"
                       style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -295px -48px;"
                       onclick="return TOOLS.action(&#39;sharpen_tool&#39;);" id="sharpen_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Burn/Dodge tool"
                       style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -497px -46px;"
                       onclick="return TOOLS.action(&#39;burn_dodge_tool&#39;);" id="burn_dodge_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Desaturate" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -547px 4px;"
                       onclick="return TOOLS.action(&#39;desaturate_tool&#39;);" id="desaturate_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Clone tool" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -346px -47px;"
                       onclick="return TOOLS.action(&#39;clone_tool&#39;);" id="clone_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Gradient" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -397px -46px;"
                       onclick="return TOOLS.action(&#39;gradient_tool&#39;);" id="gradient_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                    <a title="Crop" style="background: #989898 url(&#39;img/all.png&#39;) no-repeat -448px -48px;"
                       onclick="return TOOLS.action(&#39;crop_tool&#39;);" id="crop_tool"
                       href="https://luhuiphotoshop.netlify.app/#"></a>
                </div>
                <div style="clear:both;"></div>
                <div id="main_colour_container">
                    <input type="color" id="main_colour" value="" onchange="TOOLS.set_color(this);">
                    <div style="display:none;" id="main_colour_alt" onclick="TOOLS.toggle_color_select();"></div>
                </div>
                <div style="margin-top:10px;padding:3px 0px 3px 3px;" class="block" id="all_colors">
                    <div style="background-color:#ff0000;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ff5b31;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ffa500;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ff007f;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ff00ff;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="clear:both;"></div>
                    <div style="background-color:#00ff00;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#008000;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#7fff00;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#00ff7f;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#8ac273;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="clear:both;"></div>
                    <div style="background-color:#0000ff;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#007fff;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#37629c;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#000080;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#8000ff;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="clear:both;"></div>
                    <div style="background-color:#ffff00;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ffff80;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#ddd06a;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#808000;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#bcb88a;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="clear:both;"></div>
                    <div style="background-color:#ffffff;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#c0c0c0;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#808080;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#404040;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="background-color:#000000;" class="mini-color" onclick="TOOLS.set_color(this);"></div>
                    <div style="clear:both;"></div>
                </div>
                <div class="block">
                    <input style="width:98%;border:1px solid #393939;" type="text" id="color_hex" value="#000000"
                           onkeyup="TOOLS.set_color_manual(this)"><br>
                    <div id="main_colour_rgb">
                        Red: <input id="rgb_r" onkeyup="TOOLS.set_color_rgb(this, &#39;r&#39;)" type="text" value="255"><br>
                        Green: <input id="rgb_g" onkeyup="TOOLS.set_color_rgb(this, &#39;g&#39;)" type="text"
                                      value="255"><br>
                        Blue: <input id="rgb_b" onkeyup="TOOLS.set_color_rgb(this, &#39;b&#39;)" type="text"
                                     value="255"><br>
                        Alpha: <input id="rgb_a" onkeyup="TOOLS.set_color_rgb(this, &#39;a&#39;)" type="text"
                                      value="255">
                    </div>
                </div>
                <div class="block" id="info"><span style="font-weight:bold;min-width:45px;display:block;float:left;">Size:</span>
                    800x600<br><span style="font-weight:bold;min-width:45px;display:block;float:left;">Mouse:</span>
                    777, 763<br></div>
                <div class="block" id="action_attributes">
                    <div onclick="TOOLS.update_attribute(this, 0)" class="attribute-area" id="fill">Fill</div>
                    <div onclick="TOOLS.update_attribute(this, 0)" class="attribute-area" id="square">Square</div>
                    <div id="round_container">
                        <table style="width:100%;">
                            <tbody>
                            <tr>
                                <td style="font-weight:bold;padding-right:2px;white-space:nowrap;">Round:</td>
                                <td><input onkeyup="TOOLS.update_attribute(this);" type="text" id="round" value="0">
                                </td>
                            </tr>
                            </tbody>
                        </table>
                        <div style="float:left;width:32px;" onclick="TOOLS.update_attribute(this, 10)"
                             class="attribute-area" id="round">+
                        </div>
                        <div style="margin-left:48px;margin-bottom:15px;" onclick="TOOLS.update_attribute(this, -10)"
                             class="attribute-area" id="round">-
                        </div>
                    </div>
                </div>
            </div>
            <div id="layers_container">
                <div id="preview">
                    <canvas id="canvas_preview" width="148" height="100"></canvas>
                    <div style="margin-top:105px;padding-left:5px;">
                        <input onclick="DRAW.zoom(-1);" style="width:30px;" class="layer_add" type="button" value="-">
                        <input onclick="DRAW.zoom(+1);" style="width:30px;" class="layer_add" type="button" value="+">
                        <b>Zoom: </b><span id="zoom_nr">100</span>%
                        <br>
                        <input style="width:95%;" id="zoom_range" type="range" value="100" min="50" max="1000" step="50"
                               oninput="DRAW.zoom(this.value);">
                    </div>
                </div>
                <div id="layers_base">
                    <b>Layers</b> <a title="Add new layer" class="layer_add" onclick="LAYER.layer_add();return false;"
                                     href="https://luhuiphotoshop.netlify.app/#">+</a>
                    <a title="Move down" onclick="LAYER.move_layer(&#39;down&#39;);return false;" class="layers_arrow"
                       href="https://luhuiphotoshop.netlify.app/#">↓</a>
                    <a title="Move up" onclick="LAYER.move_layer(&#39;up&#39;);return false;" class="layers_arrow"
                       href="https://luhuiphotoshop.netlify.app/#">↑</a>
                    <a title="Set opacity" onclick="LAYER.set_alpha();return false;" class="layers_arrow"
                       href="https://luhuiphotoshop.netlify.app/#">A</a>
                    <div style="margin-top:5px;" id="layers">
                        <div class="layer active"><span class="layer_title" onclick="LAYER.select_layer(&#39;0&#39;)">Background</span><a
                                style="visibility:hidden;" class="layer_visible"
                                href="https://luhuiphotoshop.netlify.app/#"></a><a class="layer_delete" id="layer_0"
                                                                                   onclick="LAYER.layer_visibility(&#39;0&#39;);return false;"
                                                                                   title="hide"
                                                                                   href="https://luhuiphotoshop.netlify.app/#"></a>
                        </div>
                    </div>
                </div>
            </div>
            <div id="canvas_layers">
                <canvas id="canvas_back" width="800" height="600" style="width: 800px; height: 600px;">
                    <div class="error">Your browser doesn't support canvas.</div>
                </canvas>
                <canvas id="Background" width="800" height="600" style="width: 800px; height: 600px;"></canvas>
                <div id="canvas_more"></div>
                <canvas id="canvas_front" width="800" height="600" style="width: 800px; height: 600px;"></canvas>
                <canvas style="display: none; pointer-events: none; width: 800px; height: 600px;" id="canvas_grid"
                        width="800" height="600"></canvas>
            </div>
            <div id="resize-w" style="margin-left: 906px; margin-top: 301px; display: block;"></div>
            <div id="resize-h" style="margin-left: 506px; margin-top: 601px; display: block;"></div>
            <div id="resize-wh" style="margin-left: 906px; margin-top: 601px; display: block;"></div>
        </div>
        <div id="main_menu" class="ddsmoothmenu">
            <ul>
                <li style="z-index: 1000;">
                    <a  style="padding-right: 6px;">File<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;file_new&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">New</a></li>
                        <li><a onclick="MENU.do_menu([&#39;file_open&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Open...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;file_save&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Save as...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;file_print&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Print...</a></li>
                    </ul>
                </li>
                <li style="z-index: 999;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Edit<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;edit_undo&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Undo</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;edit_cut&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Cut selection</a></li>
                        <li><a onclick="MENU.do_menu([&#39;edit_copy&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Copy selection</a></li>
                        <li><a onclick="MENU.do_menu([&#39;edit_paste&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Paste selection</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;edit_select&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Select All</a></li>
                        <li><a onclick="MENU.do_menu([&#39;edit_clear&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Clear selection</a></li>
                    </ul>
                </li>
                <li style="z-index: 998;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Image<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;image_information&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Information...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_size&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Size...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_trim&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Trim</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;image_crop&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Crop Selection</a>
                        </li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;image_resize&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Resize...</a></li>
                        <li>
                            <a href="https://luhuiphotoshop.netlify.app/#">Rotate</a>
                            <ul>
                                <li><a onclick="MENU.do_menu([&#39;image_rotate_left&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Left</a></li>
                                <li><a onclick="MENU.do_menu([&#39;image_rotate_right&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Right</a></li>
                                <li><a onclick="MENU.do_menu([&#39;image_rotate&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Rotation...</a></li>
                            </ul>
                        </li>
                        <li>
                            <a href="https://luhuiphotoshop.netlify.app/#">Flip</a>
                            <ul>
                                <li><a onclick="MENU.do_menu([&#39;image_vflip&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Vertical</a></li>
                                <li><a onclick="MENU.do_menu([&#39;image_hflip&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Horizontal</a></li>
                            </ul>
                        </li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;image_colors&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Color corrections...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_auto_adjust&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Auto adjust colors</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_GrayScale&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">GrayScale</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;image_decrease_colors&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Decrease color depth...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_negative&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Negative</a></li>
                        <li><a onclick="MENU.do_menu([&#39;image_grid&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Grid...</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;image_histogram&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Histogram...</a></li>
                    </ul>
                </li>
                <li style="z-index: 997;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Layer<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;layer_new&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">New</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_dublicate&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Dublicate</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_show_hide&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Show / Hide</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_crop&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Crop Selection</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;layer_delete&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Delete</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li>
                            <a href="https://luhuiphotoshop.netlify.app/#">Move</a>
                            <ul>
                                <li><a onclick="MENU.do_menu([&#39;layer_move_up&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Up</a></li>
                                <li><a onclick="MENU.do_menu([&#39;layer_move_down&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Down</a></li>
                            </ul>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;layer_opacity&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Opacity...</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;layer_trim&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Trim</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_resize&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Resize...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_clear&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Clear</a></li>
                        <li>
                            <div class="mid-line"></div>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;layer_differences&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Differences Down</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_merge_down&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Merge Down...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;layer_flatten&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Flatten Image</a></li>
                    </ul>
                </li>
                <li style="z-index: 996;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Effects<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul id="effects_list" style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;effects_bw&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Black and White...</a>
                        </li>
                        <li>
                            <a href="https://luhuiphotoshop.netlify.app/#">Blur</a>
                            <ul>
                                <li><a onclick="MENU.do_menu([&#39;effects_BoxBlur&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Blur-Box...</a>
                                </li>
                                <li><a onclick="MENU.do_menu([&#39;effects_GaussianBlur&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Blur-Gaussian...</a>
                                </li>
                                <li><a onclick="MENU.do_menu([&#39;effects_StackBlur&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Blur-Stack...</a>
                                </li>
                                <li><a onclick="MENU.do_menu([&#39;effects_zoomblur&#39;]);"
                                       href="https://luhuiphotoshop.netlify.app/#">Blur-Zoom...</a>
                                </li>
                            </ul>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_bulge_pinch&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Bulge/Pinch...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_colorize&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Colorize...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;effects_denoise&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Denoise...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Desaturate&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Desaturate...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Dither&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Dither...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_dot_screen&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Dot Screen...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Edge&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Edge...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Emboss&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Emboss...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Enrich&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Enrich...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Gamma&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Gamma...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Grains&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Grains...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_heatmap&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Heatmap...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_HSLAdjustment&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">HSL Adjustment...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_jpg_vintage&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">JPG Compression...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Mosaic&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Mosaic...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Oil&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Oil...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_perspective&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Perspective...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Posterize&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Posterize...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Sepia&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Sepia...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Sharpen&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Sharpen...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_Solarize&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Solarize...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_tilt_shift&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Tilt Shift...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_vignette&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Vignette...</a>
                        </li>
                        <li><a onclick="MENU.do_menu([&#39;effects_vintage&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Vintage...</a>
                        </li>
                    </ul>
                </li>
                <li style="z-index: 995;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Tools<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;tools_borders&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Borders...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;tools_sprites&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Sprites...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;tools_keypoints&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Key-points</a></li>
                        <li><a onclick="MENU.do_menu([&#39;tools_color2alpha&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Color to Alpha...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;tools_color_zoom&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Color Zoom...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;tools_restore_alpha&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Restore alpha...</a></li>

                    </ul>
                </li>
                <li style="z-index: 994;">
                    <a href="https://luhuiphotoshop.netlify.app/#" style="padding-right: 6px;">Help<span
                            style="display: block;" class="rightarrowclass"></span></a>
                    <ul style="display: block; top: 30px;">
                        <li><a onclick="MENU.do_menu([&#39;help_shortcuts&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Keyboard Shortcuts...</a></li>
                        <li><a onclick="window.location=&#39;/index.htm&#39;;"
                               href="https://luhuiphotoshop.netlify.app/index.htm">Flash Photoshop</a></li>
                        <li><a onclick="window.location=&#39;/cn&#39;;" href="https://luhuiphotoshop.netlify.app/cn">简体繁体中文</a>
                        </li>
                        <li><a onclick="window.location=&#39;/zh&#39;;" href="https://luhuiphotoshop.netlify.app/zh">精简中文版</a>
                        </li>
                        <li><a onclick="window.location=&#39;/ps&#39;;" href="https://luhuiphotoshop.netlify.app/ps">多国语言</a>
                        </li>
                        <li><a onclick="window.location=&#39;/lightroom&#39;;"
                               href="https://luhuiphotoshop.netlify.app/lightroom">Lightroom CC</a></li>
                        <li><a onclick="window.location=&#39;/help&#39;;"
                               href="https://luhuiphotoshop.netlify.app/help">使用帮助</a></li>
                        <li><a onclick="MENU.do_menu([&#39;help_credits&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">Credits...</a></li>
                        <li><a onclick="MENU.do_menu([&#39;help_about&#39;]);"
                               href="https://luhuiphotoshop.netlify.app/#">About...</a></li>
                    </ul>
                </li>
            </ul>
        </div>
        <progress id="uploadprogress" min="0" max="100" value="0">0</progress>
        <div style="display:none;" id="tmp"></div>
        <div id="popup" style="top: 150px; left: 276px; display: none; overflow-y: hidden;"><h2 id="popup_drag">Save as
            ...</h2>
            <table style="width:99%;">
                <tbody>
                <tr>
                    <td style="font-weight:bold;padding-right:3px;width:130px;">File name:</td>
                    <td colspan="2"><input style="width:100%;" type="text" id="pop_data_name" value="example"
                                           placeholder="" onkeyup="POP.validate(this);"></td>
                </tr>
                <tr>
                    <td style="font-weight:bold;padding-right:3px;width:130px;">Save as type:</td>
                    <td colspan="2">
                        <div class="group"><input type="radio" checked="checked" name="type" id="pop_data_type_poptmp0"
                                                  value="PNG - Portable Network Graphics"><label
                                style="margin-right:20px;" for="pop_data_type_poptmp0">PNG - Portable Network
                            Graphics</label><br><input type="radio" name="type" id="pop_data_type_poptmp1"
                                                       value="JPG - JPG/JPEG Format"><label style="margin-right:20px;"
                                                                                            for="pop_data_type_poptmp1">JPG
                            - JPG/JPEG Format</label><br><input type="radio" name="type" id="pop_data_type_poptmp2"
                                                                value="XML - Full layers data"><label
                                style="margin-right:20px;" for="pop_data_type_poptmp2">XML - Full layers
                            data</label><br><input type="radio" name="type" id="pop_data_type_poptmp3"
                                                   value="BMP - Windows Bitmap"><label style="margin-right:20px;"
                                                                                       for="pop_data_type_poptmp3">BMP -
                            Windows Bitmap</label><br><input type="radio" name="type" id="pop_data_type_poptmp4"
                                                             value="WEBP - Weppy File Format"><label
                                style="margin-right:20px;" for="pop_data_type_poptmp4">WEBP - Weppy File
                            Format</label><br></div>
                    </td>
                </tr>
                <tr>
                    <td style="font-weight:bold;padding-right:3px;width:130px;">Quality (1-100):</td>
                    <td><input type="range" id="pop_data_quality" value="90" min="1" max="100" step="1"
                        oninput="document.getElementById(&#39;pv2&#39;).innerHTML=Math.round(this.value*100)/100;">
                    </td>
                    <td style="padding-left:10px;width:50px;" id="pv2">90</td>
                </tr>
                <tr>
                    <td style="font-weight:bold;padding-right:3px;width:130px;">Save layers:</td>
                    <td colspan="2"><input type="radio" checked="checked" name="layers" id="pop_data_layers_poptmp0"
                                           value="All"><label style="margin-right:20px;" for="pop_data_layers_poptmp0">All</label><input
                            type="radio" name="layers" id="pop_data_layers_poptmp1" value="Selected"><label
                            style="margin-right:20px;" for="pop_data_layers_poptmp1">Selected</label></td>
                </tr>
                <tr>
                    <td style="font-weight:bold;padding-right:3px;width:130px;">Trim:</td>
                    <td colspan="2"><input type="radio" checked="checked" name="trim" id="pop_data_trim_poptmp0"
                                           value="No"><label style="margin-right:20px;"
                                                             for="pop_data_trim_poptmp0">No</label><input type="radio"
                                                                                                          name="trim"
                                                                                                          id="pop_data_trim_poptmp1"
                                                                                                          value="Yes"><label
                            style="margin-right:20px;" for="pop_data_trim_poptmp1">Yes</label></td>
                </tr>
                </tbody>
            </table>
            <div style="text-align:center;margin-top:20px;margin-bottom:15px;"><input type="button"
                                                                                      onclick="POP.save();"
                                                                                      class="button" value="OK"><input
                    type="button" onclick="POP.hide();" class="button" value="Cancel"></div>
        </div>


        <div id="gg_ad"
             style="width:468px; max-height:20px;min-height:10px;position:fixed;top:5px;right:5px;z-index: 999999;">

        </div>


        <span style="display:none;">



</span>

        <div style="position:fixed;width:100%;left:5px;bottom:5px;z-index:-1"> luhuiPS online version of photo (photo)
            processing software, is a simple and practical online multi-functional photo editing tool developed using
            Flash technology. You can edit and adjust your photos here, design your pictures here, design business cards
            and other documents related to pictures. PS online image (photo) processing software functions include:
            picture scaling, rotation, copying, deleting, cropping pictures, multi-picture editing, color adjustment,
            brightness adjustment, hue adjustment, projection, glow, bevel relief, text editing, focus Blur, color
            extraction, graffiti, automatic optimization, airbrush, highlight, darken, edge blur and other functions.
            This online photo editor can edit photos online, online ps pictures, photo editing, non-mainstream photo
            editing, photo synthesis, photo production, is the best free photo editor, art photo editor. Everyone is
            welcome to use it.
        </div>


        <div class="copyright pr" style="text-align: center; margin-left: 0px;">
            <p>
                Copyright&nbsp;2004-01-11&nbsp;luhui.net&nbsp;all rights reserved
            </p>
        </div>


    </div>


</template>
<script>

    export default {
        data() {
            return {
                imgSrc: "",
                lc: null
            }
        },
        computed: {},
        watch: {},
        components: {},
        methods: {

            initData() {
                let _this = this;
                let selectedNode = _this.$store.getters.getSelectedNode
                let dirName = selectedNode.dirPath;
                let fileName = selectedNode.fileName;
                let token = localStorage.getItem("token")

            },
            injectCustomJs(jsSrc, func) {
                let _this = this;
                let head = document.getElementsByTagName('head')[0];
                let script = document.createElement('script');
                script.type = 'text/javascript';
                script.onload = function () {
                    func && func()
                }
                script.src = jsSrc;
                head.appendChild(script);
            },
            injectCustomCss(cssPath) {
                var temp = document.createElement('link');
                temp.setAttribute('rel', 'stylesheet');
                temp.setAttribute('href', cssPath);
                document.head.appendChild(temp);
            }
        },
        created() {


        },
        mounted() {
            let _this = this;
            _this.injectCustomCss("/luhuiphotoshop/styles.css")
            _this.injectCustomJs("/luhuiphotoshop/jquery-1.8.3.js", () => {
                _this.injectCustomJs("/luhuiphotoshop/menu.js", () => {
                    _this.injectCustomJs("/luhuiphotoshop/imagefilters.js", () => {
                        _this.injectCustomJs("/luhuiphotoshop/glfx.js", () => {
                            _this.injectCustomJs("/luhuiphotoshop/exif.js", () => {
                                _this.injectCustomJs("/luhuiphotoshop/exif-binaryajax.js", () => {
                                    _this.injectCustomJs("/luhuiphotoshop/harmony.js", () => {
                                        _this.injectCustomJs("/luhuiphotoshop/vintage.js", () => {
                                            _this.injectCustomJs("/luhuiphotoshop/kdtree.js", () => {
                                                _this.injectCustomJs("/luhuiphotoshop/sift.js", () => {
                                                    _this.injectCustomJs("/luhuiphotoshop/helpers.js", () => {
                                                        _this.injectCustomJs("/luhuiphotoshop/popup.js", () => {
                                                            _this.injectCustomJs("/luhuiphotoshop/controlls.js", () => {
                                                                _this.injectCustomJs("/luhuiphotoshop/tools.js", () => {
                                                                    _this.injectCustomJs("/luhuiphotoshop/draw.js", () => {
                                                                        _this.injectCustomJs("/luhuiphotoshop/layers.js", () => {
                                                                            _this.injectCustomJs("/luhuiphotoshop/menu_actions.js", () => {
                                                                                _this.injectCustomJs("/luhuiphotoshop/settings.js", () => {
                                                                                    _this.injectCustomJs("/luhuiphotoshop/main.js", () => {

                                                                                    });
                                                                                });
                                                                            });
                                                                        });
                                                                    });
                                                                });
                                                            });
                                                        });
                                                    });
                                                });
                                            });
                                        });
                                    });
                                });
                            });
                        });
                    });
                });
            });
        }
    }
</script>
