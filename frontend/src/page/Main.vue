<template>
  <div class="main">
    <div style="display: flex">
      <ui-input v-model="host" placeholder="Host" />
      <ui-button @click="getMethodList()" text="Connect" />
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
import Controller from '../component/Controller.vue';
import Axios from 'axios';

export default defineComponent({
  components: { Controller },
  async mounted() {
    this.host = localStorage.getItem('rapidebug_host') || '127.0.0.1:16000';
  },
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

      localStorage.setItem('rapidebug_host', this.host);
    },
  },
  data: () => {
    return {
      host: '127.0.0.1:16000',
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
