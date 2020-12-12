<template>
    <div id="app">
        <el-container>
            <el-header class="bg_dark">
                <div class="title">
                <img src="etcd-glyph-color.png" style="width: 22px;height: 22px;"/>
                    ETCD
                    <span>Studio</span>
                    <i>v3</i>
                </div>
            </el-header>
            <el-main class="bg_dark">
                <el-row :gutter="15">
                    <el-col :span="8">
                        <el-card class="box-card" shadow="hover" :bodyStyle="card_body_style">
                            <div slot="header" class="clearfix">
                                <el-autocomplete class="inline-input" clearable v-model="ip" :fetch-suggestions="querySearch" placeholder="请输入地址" :trigger-on-focus="true" @clear="clearIp" @select="handleSelect" @blur="handleBlur">
                                    <template slot="prepend">Http://</template>
                                </el-autocomplete>

                                <div class="l_btns">
                                    <el-tooltip v-if="false" content="切换模式" placement="bottom" effect="light">
                                        <el-button type="warning" plain circle icon="el-icon-guide" @click="changeMode"></el-button>
                                    </el-tooltip>
                                    <el-tooltip content="刷新列表" placement="bottom" effect="light">
                                        <el-button type="success" plain circle icon="el-icon-refresh" @click="refreshList"></el-button>
                                    </el-tooltip>
                                </div>
                            </div>
                            <el-tree ref="tree" :data="data" :default-expanded-keys="expand_keys" node-key="id" @node-click="handleClick" @node-contextmenu="rightHandleClick"></el-tree>
                            <ul id="menu" class="menu" v-show="menuVisible">
                                <li class="menu_item" @click="openCreateDialog('append')">添加新节点</li>
                                <li class="menu_item" v-show="allow_delete" @click="openDeleteDialog">删除节点</li>
                            </ul>
                        </el-card>
                    </el-col>
                    <el-col :span="16">
                        <el-card class="box-card" shadow="hover" :bodyStyle="card_body_style">
                            <div slot="header" class="clearfix">
                                <el-row>
                                    <el-col :span="18">
                                        <div class="node_name">{{ key_path }}</div>
                                    </el-col>
                                    <el-col :span="6">
                                        <div class="l_btns">
                                            <el-tooltip content="格式化JSON" placement="bottom" effect="light">
                                                <el-button type="warning" plain circle icon="el-icon-tickets" @click="fmtJson"></el-button>
                                            </el-tooltip>
                                            <el-tooltip content="保存内容" placement="bottom" effect="light">
                                                <el-button type="success" plain circle icon="el-icon-wallet" @click="saveContent"></el-button>
                                            </el-tooltip>
                                        </div>
                                    </el-col>
                                </el-row>
                            </div>
                            <editor ref="myEditor" @init="editorInit" :lang="lang" theme="chrome" width="100%" height="100%"></editor>
                        </el-card>
                    </el-col>
                </el-row>
            </el-main>
        </el-container>
        <el-dialog title="新建节点" width="35%" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
            <el-form ref="form" :model="form" label-position="right" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="form.name" placeholder="节点名称"></el-input>
                </el-form-item>
                <el-form-item label="TTL">
                    <el-input v-model="form.ttl" type="number" min=1 value="" placeholder="超时时间，单位秒，不填则默认为永不超时"></el-input>
                </el-form-item>
                <el-form-item label="值">
                    <el-input type="textarea" :rows="5" v-model="form.value" placeholder="节点值"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="createNode">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import axios from "axios";
export default {
    name: "app",
    data() {
        return {
            api: "",
            mode: "tree",
            data: [],
            expand_keys: [],
            lang: "text",
            key_path: "",
            menuVisible: false,
            allow_delete: true,
            card_body_style: {
                padding: "0px",
                height: "calc(100vh - 157px)",
                overflow: "auto"
            },
            ip: "",
            ip_attrs: [],
            dialogFormVisible: false,
            active_node: {},
            form: {
                name: '',
                ttl: '',
                dir: true,
                value: ''
            }
        };
    },
    methods: {
        editorInit: function () {
            require("brace/ext/language_tools"); //language extension prerequsite...
            require("brace/mode/html");
            require("brace/mode/javascript"); //language
            require("brace/mode/json"); //language
            require("brace/mode/less");
            require("brace/theme/chrome");
            require("brace/snippets/javascript"); //snippet

            this.editor = this.$refs.myEditor.editor;
            this.editor.getSession().setTabSize(4);
            // this.editor.setReadOnly(true); 
            this.editor.resize();
        },
        fmtJson() {
            this.format("json");
        },
        format(type) {
            if (type === "json") {
                let val = JSON.parse(this.editor.getValue());
                this.lang = 'json';
                this.editor.setValue(JSON.stringify(val, null, 4));
            }
            this.editor.clearSelection();
            this.editor.navigateFileStart();
        },
        handleClick(tab) {
            this.active_node = tab;
            this.menuVisible = false;

            if (this.ip == '') return;
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });

            axios.get(this.api + "/api/v3/value", {
                params: {
                    endpoint: this.ip,
                    key: this.active_node.id
                }
            }
            ).then((response) => {
                loading.close();
                if (response.data.error) {
                    this.$message({
                        type: 'info',
                        message: response.data.error
                    });
                    return;
                }
                this.lang = 'text';
                this.editor.setValue(response.data.value);
                this.editor.clearSelection();
                this.key_path = tab.id;
            }).catch(function (error) {
                console.log(error);
                loading.close();
            });
        },
        rightHandleClick(event, tab) {
            event.preventDefault();
            this.active_node = tab;
            this.allow_delete = (tab.id == '/') ? false : true;
            this.menuVisible = false; // 先把模态框关死，目的是 第二次或者第n次右键鼠标的时候 它默认的是true
            this.menuVisible = true; // 显示模态窗口，跳出自定义菜单栏
            var menu = document.querySelector("#menu");
            document.addEventListener("click", this.foo); // 给整个document添加监听鼠标事件，点击任何位置执行foo方法
            menu.style.display = "block";
            menu.style.left = event.clientX - 0 + "px";
            menu.style.top = event.clientY - 20 + "px";
        },
        foo() {
            // 取消鼠标监听事件 菜单栏
            this.menuVisible = false;
            document.removeEventListener("click", this.foo); // 要及时关掉监听，不关掉的是一个坑，不信你试试，虽然前台显示的时候没有啥毛病，加一个alert你就知道了
        },
        querySearch(queryString, cb) {
            let ips = this.ip_attrs;
            let results = queryString
                ? ips.filter(this.createFilter(queryString))
                : ips;
            // 调用 callback 返回建议列表的数据
            cb(results);
        },
        clearIp() {
            this.data = [];
            this.mode = 'tree';
        },
        refreshList() {
            if (this.ip == "") return;
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });

            let trie = true;
            if (this.mode == 'tree') trie = true;
            if (this.mode == 'ties') trie = false;

            axios.get(this.api + "/api/v3/keys", {
                params: {
                    endpoint: this.ip,
                    trie: trie,
                    refresh: true
                }
            }).then((response) => {
                loading.close();
                if (response.data.error) {
                    this.$message({
                        type: 'error',
                        message: '错误信息：' + response.data.error
                    });
                    return;
                }
                if (this.mode == 'tree') {
                    this.data = [response.data]
                    this.expand_keys.push(response.data.id);
                }

                if (this.mode == 'ties') {
                    let data = this.formatArr(response.data);
                    this.data = data;
                }
                this.$message({
                    type: 'success',
                    message: '节点列表加载完成'
                });

                let is_include = this.ip_attrs.some(item => {
                    if (item.value == this.ip) return true;
                });
                if (is_include) return;
                this.ip_attrs.push({
                    value: this.ip
                });
                localStorage.setItem("attrs", JSON.stringify(this.ip_attrs));
            }).catch(function (error) {
                console.log(error);
                loading.close();
            });
        },
        changeMode() {
            if (this.mode == 'tree') {
                this.mode = 'ties'
            } else if (this.mode == 'ties') {
                this.mode = 'tree'
            }
            this.refreshList();
        },
        formatArr(arr) {
            arr = arr.map(item => {
                return {
                    id: item,
                    label: item
                }
            })
            return arr;
        },
        createFilter(queryString) {
            return ip => {
                return (
                    ip.value
                        .toLowerCase()
                        .indexOf(queryString.toLowerCase()) === 0
                );
            };
        },
        handleBlur() {
            if (this.ip == "") return;
        },
        handleSelect(item) {
            if (item.value == "") return;
            this.ip = item.value;
            this.refreshList();
        },
        openCreateDialog(type) {
            this.add_mode = type;
            this.dialogFormVisible = true;
        },
        createNode() {
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });

            let key = '';
            if (this.active_node.id == '/') {
                key = '/' + this.form.name;
            } else {
                key = this.active_node.id + "/" + this.form.name
            }

            axios.post(this.api + "/api/v3/key?endpoint=" + this.ip, {
                key: key,
                value: this.form.value,
                ttl: this.form.ttl ? Number(this.form.ttl) : -1
            }
            ).then((response) => {
                loading.close();
                if (response.data.error) {
                    this.$message({
                        type: 'info',
                        message: response.data.error
                    });
                    return;
                }
                let newNode = {
                    id: key,
                    label: this.form.name,
                    children: []
                }
                this.$refs.tree.append(newNode, this.active_node)
            }).catch(function (error) {
                console.log(error);
                loading.close();
            });

            this.dialogFormVisible = false;
        },
        openDeleteDialog() {
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });
            let id = this.active_node.id;
            let name = this.active_node.label;
            this.$confirm(`删除 "${name}", 是否继续?`, '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {

                axios.delete(this.api + "/api/v3/key?endpoint=" + this.ip + "&key=" + id,
                ).then((response) => {
                    loading.close();
                    if (response.data.error) {
                        this.$message({
                            type: 'info',
                            message: response.data.error
                        });
                        return;
                    }
                    this.$refs.tree.remove(this.active_node)
                    this.$message({
                        type: 'success',
                        message: '删除成功!'
                    });
                }).catch(function (error) {
                    console.log(error);
                    loading.close();
                });
            }).catch(() => {
                loading.close();
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            });
        },
        saveContent() {
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });
            let content = this.editor.getValue();
            axios.put(this.api + "/api/v3/value?endpoint=" + this.ip, {
                "key": this.active_node.id,
                "value": content
            }
            ).then((response) => {
                loading.close();
                if (response.data.error) {
                    this.$message({
                        type: 'info',
                        message: response.data.error
                    });
                    return;
                }
                this.$message({
                    type: 'success',
                    message: '保存成功!'
                });
            }).catch(function (error) {
                console.log(error);
                loading.close();
            });
        }
    },
    mounted() {
        let local_ip_attrs = localStorage.getItem("attrs");
        this.ip_attrs = local_ip_attrs ? JSON.parse(local_ip_attrs) : [];
        
        /*global PLATFROM_CONFIG*/
        this.api = PLATFROM_CONFIG.BASE_URL;
    },
    components: {
        editor: require("vue2-ace-editor")
    }
};
</script>

<style>
body {
    padding: 0px;
    margin: 0px;
}

.bg_dark {
    background: rgb(84, 92, 100);
}

.title {
    font-size: 1.5em;
    font-weight: 700;
    color: #fff;
    text-align: left;
    margin-top: 38px;
}

.title > span {
    color: rgb(255, 208, 75);
}

.title > span + i {
    padding-left: 10px;
    font-style: normal;
    font-size: 18px;
    color: rgb(255, 208, 75);
}

#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    /* text-align: center; */
    color: #2c3e50;
}

.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
}

.el-card__header {
    padding: 7px !important;
    text-align: left;
    height: 55px;
    overflow: hidden;
}

.l_btns {
    float: right;
    padding: 0 3px;
}

.node_name {
    height: 40px;
    line-height: 40px;
    padding-left: 4px;
    font-weight: bold;
    color: #777;
}

.ace-chrome {
    font-size: 13px;
}

.menu {
    position: fixed;
    border: 1px solid #e4e7ed;
    border-radius: 4px;
    background-color: #fff;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    box-sizing: border-box;
    font-size: 14px;
    list-style: none;
    padding: 0px;
    min-width: 110px;
}

.menu > li {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    box-sizing: border-box;
    padding: 0 20px;
    height: 34px;
    line-height: 34px;
    cursor: pointer;
    color: #606266;
}

.menu > li:hover {
    background: #f5f7fa;
}

#app .el-tree-node > .el-tree-node__children {
    overflow: visible;
}

.el-message-box {
    vertical-align: text-bottom !important;
}
</style>
