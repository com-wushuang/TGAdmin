<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="时间间隔（s）:"><el-input v-model.number="formData.interval" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="active字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.active" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="advertisement字段:">
                <el-input v-model="formData.advertisement" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="expire字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.expire" clearable></el-date-picker>
           </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createJob,
    updateJob,
    findJob
} from "@/api/job";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Job",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            interval:0,
            active:false,
            advertisement:"",
            expire:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createJob(this.formData);
          break;
        case "update":
          res = await updateJob(this.formData);
          break;
        default:
          res = await createJob(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findJob({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rejob
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>