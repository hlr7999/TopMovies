<template>
<el-container direction="vertical">
  <el-main>
    <el-card
      class="box-card"
      v-loading="loadingAll"
      element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading">

      <div slot="header" class="clearfix">
        <div class="title">
          <span>{{film.title}}</span>
          <span class="year">({{film.year}})</span>
        </div>
        <el-button
          type="text"
          class="home"
          @click="goHome">首页</el-button>
      </div>

      <el-row class="mainInfo">
        <el-col style="width:140px">
          <el-image :src="film.poster" class="poster"></el-image>
        </el-col>
        <el-col :span="10" class="info">
          <div class="itemTitle">
            <span style="color:gray;">导演:</span>
            <span v-for="(item, index) in film.directors" :key="index" style="color:#37a">
              {{item.name}}<span v-if="index < film.directors.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">编剧:</span>
            <span v-for="(item, index) in film.writers" :key="index" style="color:#37a">
              {{item.name}}<span v-if="index < film.writers.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">演员:</span>
            <span v-for="(item, index) in film.casts" :key="index" style="color:#37a">
              {{item.name}}<span v-if="index < film.casts.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">类型:</span>
            <span v-for="(item, index) in film.genres" :key="index" style="color:#37a">
              {{item}}<span v-if="index < film.genres.length - 1"> / </span>
            </span>            
          </div>
          <div class="itemTitle">
            <span style="color:gray;">制片国家/地区:</span>
            <span v-for="(item, index) in film.countries" :key="index" style="color:#37a">
              {{item}}<span v-if="index < film.countries.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">语言:</span>
            <span v-for="(item, index) in film.languages" :key="index" style="color:#37a">
              {{item}}<span v-if="index < film.languages.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">上映日期:</span>
            <span v-for="(item, index) in film.pubdate" :key="index" style="color:#37a">
              {{item}}<span v-if="index < film.pubdate.length - 1"> / </span>
            </span>
          </div>
          <div class="itemTitle">
            <span style="color:gray;">片长:</span>
            <span style="color:#37a">{{film.duration}}分钟</span>
          </div>
        </el-col>
        <el-col :span="8" style="height:190px">
          <el-row class="tRate">
            <span>总评({{film.rating.rating_people}}人评价)</span>
            <el-row>
              <el-col :span="16">
                <el-rate
                  v-model="film.rating_star"
                  disabled
                  text-color="#ff9900"
                  disabled-void-color="#dddddd"
                  style="float:right">
                </el-rate>
              </el-col>
              <el-col :span="8">
                <span class="score">{{film.rating.average}}</span>
              </el-col>
            </el-row>
          </el-row>
          <el-row class="myRate">
            <span>你的评分</span>
            <el-rate 
              allow-half="true"
              v-model="youRate">
            </el-rate>
          </el-row>
        </el-col>
      </el-row>

      <el-row class="summary">
        <div class="summaryTitle">剧情简介......</div>
        <p class="summaryContent">
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{film.summary}}
        </p>
      </el-row>

    </el-card>
  </el-main>
</el-container>
</template>

<script>
export default {
  data() {
    return {
      film: {},
      youRate: 0,
      loadingAll: true
    }
  },
  methods: {
    getData() {
      let filmid = this.$route.params.id
      this.$axios.get(`http://47.100.225.59:10080/films/${filmid}`)
      .then(res => {
        this.film = res.data
        this.film.rating_star = Number(this.film.rating.average) / 2
        this.loadingAll = false
      })
    },
    goHome() {
      this.$router.push("/")
    }
  },
  mounted() {
    this.getData()
  }
}
</script>

<style scoped>
.title {
  font-size: 23px;
  float: left;
  padding: 3px 0;
}

.year {
  color: gray;
}

.home {
  font-size: 23px;
  float: right;
  padding: 3px 10px;
  margin-left: 66px;
  color: black;
}

.home:hover {
  color: gray;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both
}

.box-card {
  width: 888px;
  margin: auto;
}

.mainInfo {
  border-bottom: 1px solid #EBEEF5;
  margin-bottom: 50px;
  min-height: 210px;
}

.poster{
  width: 135px;
  height: 191px;
  float: left;
}

.info {
  min-height: 190px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: flex-start;
  border-right: 1px solid #EBEEF5;
}

.el-col .itemTitle {
  font-size: 13px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}

.tRate {
  height:50%; 
  border-bottom:1px solid #EBEEF5;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.myRate {
  height: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.score {
  color: rgb(247, 168, 42);
  font-size: 17px;
  margin-left: 5px;
  float: left;
}

.summary {
  margin-top: 10px;
}

.summaryTitle {
  text-align: left;
  margin-bottom: 10px;
  color: green;
  margin-left: 10px;
  font-size: 20px;
}

.summaryContent {
  text-align: left;
  line-height: 30px;
  margin: 0 15px;
}
</style>
