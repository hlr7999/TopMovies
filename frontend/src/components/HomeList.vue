<template>
<el-container direction="vertical">
  <el-main>
    <el-card class="box-card">

      <div slot="header" class="clearfix">
        <span class="title">电影排行榜</span>
        <div style="float: right; padding: 3px 0">
          <span style="color: gray; font-size: 13px">排序方式</span>
          <el-select 
            v-model="sortWay" 
            placeholder="选择排序方式" 
            size="mini" 
            @change="changeSort">
            <el-option
              v-for="item in sortWaies"
              :key="item.id"
              :label="item.label"
              :value="item.id">
            </el-option>
          </el-select>
        </div>
      </div>

      <el-table
        v-loading="loadingAll"
        element-loading-text="拼命加载中"
        element-loading-spinner="el-icon-loading"
        :data="currentFilmList">
        <el-table-column label="" width="80">
          <template slot-scope="scope">
            <el-image :src="scope.row.poster" class="poster"></el-image>
          </template>
        </el-table-column>
        <el-table-column label="电影名">
          <template slot-scope="scope">
            <span>{{ scope.row.title }}<br/>{{scope.row.pubdate}}</span>
          </template>
        </el-table-column>
        <el-table-column label="评分">
          <template slot-scope="scope">
            <el-rate
              v-model="scope.row.rating_star"
              disabled
              text-color="#ff9900"
              disabled-void-color="#dddddd">
            </el-rate>
            <span class="score">{{scope.row.rating.average}}</span>
            <br/>
            <span>{{scope.row.rating.rating_people}}人评价</span>
          </template>
        </el-table-column>
        <el-table-column label="简介" width="">
          <template slot-scope="scope">
            <span v-if="scope.row.summary.length > 50">
              {{scope.row.summary.slice(0,50)}}...
            </span>
            <span v-else>
              {{scope.row.summary}}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="">
          <template slot-scope="scope">
            <el-button
              type="text"
              @click="viewIntro(scope.row._id)"
              style="color:blue; margin-left:40px">详情</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <el-pagination
        layout="total, prev, pager, next"
        :page-size="10"
        :current-page.sync="currentPage"
        @current-change="changePage"
        style="margin-top: 10px"
        :total="filmList.length">
      </el-pagination>

    </el-card>
  </el-main>
</el-container>
</template>

<script>
export default {
  data() {
    return {
      loadingAll: true,
      filmList: this.$store.state.filmList,
      currentFilmList: [],
      sortWaies: [{
        id: 0,
        label: "评分"
      }, {
        id: 1,
        label: "评价人数"
      }, {
        id: 2,
        label: "上映时间"
      }],
      sortWay: this.$store.state.sortWay,
      currentPage: this.$store.state.currentPage
    }
  },
  methods: {
    getData() {
      if (!this.filmList.length) {
        this.$axios.get("http://47.100.225.59:10080/films")
        .then(res => {
          this.filmList = res.data
          this.filmList.forEach(item => {
            item.rating_star = Number(item.rating.average) / 2
            item.pubdate = item.pubdate[0].split("(")[0]
          })
          this.$store.commit("newFilmList", this.filmList)
          this.loadingAll = false
          this.currentFilmList = this.filmList.slice(
            (this.currentPage - 1) * 10,
            (this.currentPage - 1) * 10 + 10
          )
        })
      } else {
        this.loadingAll = false
        this.currentFilmList = this.filmList.slice(
          (this.currentPage - 1) * 10,
          (this.currentPage - 1) * 10 + 10
        )
      }
    },
    changeSort() {
      this.loadingAll = true
      switch (this.sortWay) {
      case 0:
        var compare = function(x, y) {
          var valx = x.rating.average
          var valy = y.rating.average
          if (valx > valy) {
            return -1
          } else if (valx < valy) {
            return 1
          } else {
            return 0
          }
        }
        this.filmList.sort(compare)
        break;
      case 1:
        var compare = function(x, y) {
          var valx = Number(x.rating.rating_people)
          var valy = Number(y.rating.rating_people)
          if (valx > valy) {
            return -1
          } else if (valx < valy) {
            return 1
          } else {
            return 0
          }
        }
        this.filmList.sort(compare)
        break;
      case 2:
        var compare = function(x, y) {
          var valx = x.pubdate
          var valy = y.pubdate
          if (valx > valy) {
            return -1
          } else if (valx < valy) {
            return 1
          } else {
            return 0
          }
        }
        this.filmList.sort(compare)
        break;
      }
      this.loadingAll = false
      this.currentPage = 1
      this.$store.commit("newPage", this.currentPage)
      this.$store.commit("newSort", this.sortWay)
      this.currentFilmList = this.filmList.slice(0, 10)
    },
    changePage: function() {
      this.currentFilmList = this.filmList.slice(
        (this.currentPage - 1) * 10,
        (this.currentPage - 1) * 10 + 10
      )
      this.$store.commit("newPage", this.currentPage)
    },
    viewIntro: function(id) {
      this.$router.push(`/description/${id}`)
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

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both
}

.box-card {
  width: 999px;
  margin: auto;
}

.poster {
  width: 67px;
  height: 100px;
}

.el-rate {
  float: left;
}

.score {
  color: rgb(247, 168, 42);
  font-size: 17px;
  margin-left: 5px;
}
</style>
