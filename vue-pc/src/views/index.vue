<template>
  <div class="page-container">
    <div class="sidebar-menu toggle-others fixed">
      <div class="sidebar-menu-inner">
        <header class="logo-env">
          <!-- logo -->
          <div class="logo">
            <a href="#" class="logo-expanded">
              <img src="../assets/images/logo@2x.png" width="100%" alt="" />
            </a>
            <a href="#" class="logo-collapsed">
              <img
                src="../assets/images/logo-collapsed@2x.png"
                width="40"
                alt=""
              />
            </a>
          </div>
          <div class="mobile-menu-toggle visible-xs">
            <a href="#" data-toggle="user-info-menu">
              <i class="linecons-cog"></i>
            </a>
            <a href="#" data-toggle="mobile-menu">
              <i class="fa-bars"></i>
            </a>
          </div>
        </header>
        <!-- 侧边栏 -->
        <ul id="main-menu" class="main-menu">
          <li v-for="(menu, idx) in items" :key="idx">
            <a :href="'#' + transName(menu)" class="smooth">
              <i :class="menu.cover"></i>
              <span class="title">{{ transName(menu) }}</span>
              <!-- <span class="label label-Primary pull-right hidden-collapsed">编辑</span> -->
            </a>
          </li>
          <!-- manage -->
          <!-- <li>
              <a href="about.html">
                  <i class="linecons-heart"></i>
                  <span class="tooltip-blue">添加分类</span>
              </a>
          </li>
          <li>
              <a href="about.html">
                  <i class="linecons-heart"></i>
                  <span class="tooltip-blue">添加收录</span>
              </a>
          </li> -->
        </ul>
      </div>
    </div>

    <div class="main-content">
      <nav class="navbar user-info-navbar" role="navigation">
        <ul class="user-info-menu left-links list-inline list-unstyled">
          <li class="hidden-sm hidden-xs">
            <a href="#" data-toggle="sidebar"><i class="fa-bars"></i></a>
          </li>
          <!-- <li class="dropdown hover-line language-switcher">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <img :src="lang.flag" /> {{ lang.name }}
            </a>
            <ul class="dropdown-menu languages">
              <li
                :class="{ active: langItem.key === lang.key }"
                v-for="langItem in langList"
                :key="langItem.key"
              >
                <a href="#" @click="lang = langItem">
                  <img :src="langItem.flag" /> {{ langItem.name }}
                </a>
              </li>
            </ul>
          </li> -->
        </ul>
        <ul class="user-info-menu right-links list-inline list-unstyled">
            <li>
                <a href="javascript:void(0);">
                    登录
                </a>
            </li>
            <li>
                <a href="javascript:void(0);" @click="dialogFormRegister = true">
                    注册
                </a>
            </li>
            <!-- <li class="hidden-sm hidden-xs">
                <a href="#" data-toggle="modal" data-target="#register" target="_blank">
                    <i class="fa-github"></i>  GitHub
                </a>
            </li> -->
        </ul>
      </nav>

      <div v-for="(item, idx) in items" :key="idx">
        <div v-if="item.collect">
          <WebItem :item="item" :transName="transName" />
        </div>
        <!-- <div v-else v-for="(subItem, idx) in item.children" :key="idx">
          <WebItem :item="subItem" :transName="transName" />
        </div> -->
      </div>

      <Footer />
      <!-- register -->
      <el-dialog title="注册" :visible.sync="dialogFormRegister">
          <el-form :model="form">
              <el-form-item label="手机号" :label-width="formLabelWidth">
                  <el-input v-model="registerForm.phone" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="验证码" :label-width="formLabelWidth">
                  <el-input v-model="registerForm.code" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="密码" :label-width="formLabelWidth">
                  <el-input v-model="registerForm.password" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="重复密码" :label-width="formLabelWidth">
                  <el-input v-model="registerForm.confirm_password" autocomplete="off"></el-input>
              </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormRegister = false">取 消</el-button>
              <el-button type="primary" @click="dialogFormRegister = false">确 定</el-button>
          </div>
      </el-dialog>
    </div>

  </div>
</template>

<script>
import WebItem from "../components/WebItem.vue";
import Footer from "../components/Footer.vue";
import {getHome} from "../api/index"
import { loadJs } from '../assets/js/app.js'

export default {
  name: "Index",
  components: {
    WebItem,
    Footer,
  },
  data() {
    return {
      items: [],
      collectCategoryModelList: [],
      userInfo: [],
      lang: {},
      langList: [
        {
          key: "zh",
          name: "简体中文",
          flag: "./assets/images/flags/flag-cn.png",
        },
        {
          key: "en",
          name: "English",
          flag: "./assets/images/flags/flag-us.png",
        },
      ],
      dialogFormRegister: false,
      registerForm: [],
    };
  },
  created() {
    this.lang = this.langList[0];
    loadJs();
  },
  methods: {
    getHome() {
        let ajaxData = {};
        getHome(ajaxData).then(res => {
            this.collectCategoryModelList = res.data.collectCategoryModelList || []
            this.userInfo = res.data.userInfo || []
            this.items = this.collectCategoryModelList
        }).catch(err => {
            console.log(err)
        })
    },
    transName(webItem) {
      return webItem.name;
      // return this.lang.key === "en" ? webItem.en_name : webItem.name;
    },
  },
   mounted() {
      this.getHome();
  }
};
</script>

<style>
</style>
