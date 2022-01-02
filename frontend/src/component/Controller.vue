<template>
  <div :class="$style.controller">
    <div :class="$style.header">
      <div>{{ controller.name }}</div>
      <ui-button
        @click="isShow = !isShow"
        icon="arrow_down_small"
        style="flex: 0; padding: 2px; border-radius: 0; margin-left: auto"
      />
    </div>
    <div v-if="isShow" :class="$style.body">
      <Method
        v-for="x in controller.methodList"
        :method="x"
        :host="host"
        :key="x.httpMethod + '' + x.fullPath"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Method from './Method.vue';

export default defineComponent({
  props: {
    controller: Object,
    host: {
      type: String,
      required: true,
    },
  },
  components: { Method },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      isShow: false,
    };
  },
});
</script>

<style module lang="scss">
.controller {
  margin-bottom: 5px;
  padding: 10px;
  background: #232323;

  .header {
    display: flex;
    align-items: center;

    > div {
      background: #156d04;
      padding: 5px 15px;
      border-radius: 16px;
      font-weight: bold;
      margin-right: 10px;
      user-select: none;
      color: #87ff8b;
      text-transform: uppercase;
      font-size: 14px;
    }

    img {
      display: block;
      margin-left: auto;
      background: #2b2b2b;
      padding: 10px;
      border-radius: 4px;
    }
  }

  .body {
    margin-top: 10px;
  }
}
</style>
