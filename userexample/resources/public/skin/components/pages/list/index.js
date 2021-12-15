define(function(require) {
  var itemlist = require("js/app/apis/user/list");
  var itemenable = require("js/app/apis/user/enable");
  var itemdisable = require("js/app/apis/user/disable");

  var lodash = require("lodash");
  return {
    name: "compomentname",
    watch: {
      $route: function(to, from) {
        this.load();
      }
    },
    mounted: function() {
      this.load();
    },
    methods: {
      load: function() {
        var self = this;
        this.Items=[]
        this.Last = this.$route.query.last;
        this.Rev = this.$route.query.rev?true:false;
        itemlist(self, function() {});
      },
      prev:function(){
        var query = lodash.clone(this.$route.query);
        query.last=this.Items[0].ID
        query.rev="true"
        this.$router.push({ query: query });
      },
      next:function(){
        var query = lodash.clone(this.$route.query);
        query.last=this.Items[this.Items.length-1].ID
        query.rev=""
        this.$router.push({ query: query });
      },
      handleEnable: function(item) {
        var self = this;
        self.errors = [];
        this.$confirm("此操作将激活用户, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(function() {
            itemenable(self, item.ID, function() {
              self.load();
            });
          })
          .catch(function() {});
      },
      handleDisable: function(item) {
        var self = this;
        self.errors = [];
        this.$confirm("此操作将禁用用户, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(function() {
            itemdisable(self, item.ID, function() {
              self.load();
            });
          })
          .catch(function() {});
      },
    },
    template: require("text!./index.html"),
    data: function() {
      return {
        Items: [],
        Sort: "",
        Rev: false,
        Last:"",
        Iter:"",
        errors: [],
      };
    }
  };
});
