<template>
  <div :class="$style.method">
    <div :class="$style.header">
      <div :class="[$style.type, $style[method.type]]">{{ method.type }}</div>
      <div>{{ method.fullPath }}</div>
      <img @click="isShow = !isShow" class="clickable" src="../asset/arrow_down.svg" alt="" />
    </div>
    <div v-if="isShow" :class="$style.body">
      <div v-for="x in method.struct" :key="x.name">{{ x.name }} {{ x.kind }} {{ x.value }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    method: Object,
  },
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
.method {
  margin-bottom: 5px;

  &:last-child {
    margin-bottom: 0;
  }

  .header {
    display: flex;
    align-items: center;
    padding: 10px;
    background: #424242;
    border-left: 5px solid #6b6b6b;
    box-sizing: border-box;

    img {
      display: block;
      margin-left: auto;
      background: #2b2b2b;
      padding: 10px;
      border-radius: 4px;
      user-select: none;
    }

    .type {
      background: #666666;
      padding: 5px 12px;
      border-radius: 4px;
      font-weight: bold;
      margin-right: 10px;
      user-select: none;
      font-size: 14px;
    }

    .GET {
      background: #007dfe;
    }

    .POST {
      background: #fe9000;
    }

    .PATCH {
      background: #07ac1d;
    }

    .DELETE {
      background: #e0250c;
    }
  }
}
</style>
