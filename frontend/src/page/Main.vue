<template>
  <div class="main">
    <div style="display: flex">
      <Input v-model="host" placeholder="Host" />
      <Button @click="getMethodList()" text="Connect" />
    </div>

    <div class="body">
      <div v-for="x in controllerList" :key="x.name">
        <Controller :controller="x" :host="host" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Input from '../component/Input.vue';
import Button from '../component/Button.vue';
import Controller from '../component/Controller.vue';
import Axios from 'axios';

export default defineComponent({
  components: { Input, Button, Controller },
  async mounted() {},
  methods: {
    async getMethodList() {
      const methodList = (await Axios.get(`http://${this.host}/debug/api/methodList`)).data
        .response;
      methodList.sort((a: any, b: any) => {
        return a.controller.localeCompare(b.controller) || b.httpMethod.localeCompare(a.httpMethod);
      });

      const controller = {} as any;
      for (let i = 0; i < methodList.length; i++) {
        if (!controller[methodList[i].controller]) {
          controller[methodList[i].controller] = {
            name: methodList[i].controller,
            methodList: [], //s
          };
        }

        controller[methodList[i].controller].methodList.push(methodList[i]);
      }

      this.controllerList.length = 0;
      for (let s in controller) {
        this.controllerList.push(controller[s]);
      }
    },
  },
  data: () => {
    return {
      host: '127.0.0.1:5000',
      controllerList: [] as any[],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  .body {
    padding: 10px;
  }
}
</style>
