<template>
  <div :class="$style.struct">
    <div :class="$style.header">
      <div :class="$style.name">{{ struct.name }}</div>
      <div :class="$style.type">{{ struct.kind }}</div>
      <input type="text" v-if="struct.kind === 'string'" v-model="obj[struct.name]" />
      <input
        type="number"
        v-if="struct.kind.match('int') || struct.kind.match('float')"
        v-model="obj[struct.name]"
      />
      <input type="checkbox" v-if="struct.kind === 'bool'" v-model="obj[struct.name]" />
      <input type="datetime-local" v-if="struct.type === 'time.Time'" v-model="obj[struct.name]" />
    </div>
    <div v-if="struct.type !== 'time.Time'">
      <div v-for="x in struct.fieldList" :key="x.name">
        <Struct
          v-if="x.kind === 'struct' && x.type !== 'time.Time'"
          :struct="x"
          :obj="obj[x.name]"
        />
        <Struct v-else :struct="x" :obj="obj" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    struct: Object,
    obj: [Object, String, Array, Number, Boolean],
  },
  components: {},
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
.struct {
  margin-bottom: 5px;
  padding-left: 25px;

  .header {
    display: flex;
    align-items: center;

    .name {
      margin-right: 10px;
      color: #e7e7e7;
    }

    .type {
      color: #fe6600;
      font-weight: bold;
    }

    input {
      background: #1b1b1b;
      border: none;
      padding: 5px;
      outline: none;
      margin-left: 10px;
      border-radius: 4px;
      color: #d8d8d8;
      width: 100%;
      box-sizing: border-box;
    }
  }
}
</style>
