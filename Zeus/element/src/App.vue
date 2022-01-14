<template>
  <div id="app">
    <el-table
      :data="tableData"
      style="width: 100%">
      <!-- el-table-column标签表示一个数据项（表格列），prop是数据项字段名，label为数据项展示名称 -->
      <el-table-column
        prop="HX_ID"
        label="日期"
        width="180">
      </el-table-column>
      <el-table-column
        prop="HX_NAME"
        label="姓名"
        width="180">
      </el-table-column>
      <el-table-column
        prop="address"
        label="地址">
      </el-table-column>
    </el-table>
    <div class="tabListPage">
           <el-pagination @size-change="handleSizeChange" 
                          @current-change="handleCurrentChange" 
                          :current-page="currentPage" 
                          :page-sizes="pageSizes" 
                          :page-size="PageSize" layout="total, sizes, prev, pager, next, jumper" 
                          :total="totalCount">
             </el-pagination>
       </div>
  </div>
</template>

<script>

export default {
  name: "App",
  components: {},
  data() {
    return {
      tableData: [],
      pageSizes: 10,
      currentPage: 1
    }
  },
  mounted() {
    // 请求接口数据
    fetch('http://127.0.0.1:9001/v1/Curve/GetAllCurve').then(result => {
      return result.json();
    }).then(result => {
      const data = result.data;
      console.log(data)
      console.log(data.length)
      // 赋值给tableData
      this.tableData = data
      this.pageSizes = data.length / 20;
      this.curpage = 1;
    }).catch(e => {
      console.log(e);
      // mock
      
    })
  },
  method() {
  
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
