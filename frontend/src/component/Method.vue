<template>
  <div :class="$style.method">
    <div :class="$style.header">
      <div :class="[$style.type, $style[method.httpMethod]]">{{ method.httpMethod }}</div>
      <div>{{ method.fullPath }}</div>
      <img @click="isShow = !isShow" class="clickable" src="../asset/arrow_down.svg" alt="" />
    </div>
    <div v-if="isShow && isReady" :class="$style.body">
      <Struct :struct="method.input" :obj="obj" style="padding: 0" />
      <div>
        <pre v-html="formatHighlight(JSON.stringify(obj, null, 4), colors)"></pre>
      </div>
      <button @click="execute()">Execute</button>
      <div>
        <pre v-html="formatHighlight(JSON.stringify(response, null, 4), colors)"></pre>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Struct from './Struct.vue';
// @ts-ignore
import formatHighlight from 'json-format-highlight';
import Moment from 'moment';
import Axios from 'axios';

export default defineComponent({
  props: {
    method: {
      type: Object,
      required: true,
    },
    host: {
      type: String,
      required: true,
    },
  },
  components: { Struct },
  async mounted() {
    this.obj = this.prepareStruct(this.method.input, {});
    this.isReady = true;
  },
  methods: {
    formatHighlight,
    prepareStruct(struct: any, out: any) {
      if (struct.kind === 'struct') {
        for (let i = 0; i < struct.fieldList.length; i++) {
          if (struct.fieldList[i].kind === 'string') {
            out[struct.fieldList[i].name] = '';
          }
          if (struct.fieldList[i].kind === 'bool') {
            out[struct.fieldList[i].name] = false;
          }
          if (struct.fieldList[i].kind.match('int') || struct.fieldList[i].kind.match('float')) {
            out[struct.fieldList[i].name] = 0;
          }
          if (struct.fieldList[i].kind === 'slice') {
            out[struct.fieldList[i].name] = [];
          }
          if (struct.fieldList[i].kind === 'map') {
            out[struct.fieldList[i].name] = {};
          }
          if (struct.fieldList[i].kind === 'struct') {
            if (struct.fieldList[i].type === 'time.Time') {
              out[struct.fieldList[i].name] = Moment().format('YYYY-MM-DD HH:mm:ss');
            } else {
              out[struct.fieldList[i].name] = this.prepareStruct(struct.fieldList[i], {});
            }
          }
        }
      }
      return out;
    },
    async execute() {
      if (this.method.httpMethod === 'POST') {
        try {
          this.response = (
            await Axios.post(`http://${this.host}${this.method.fullPath}`, this.obj, {
              headers: {
                'Content-Type': 'application/json',
              },
            })
          ).data;
        } catch (e) {
          this.response = e.response.data;
        }
      }
    },
  },
  data: () => {
    return {
      isShow: false,
      obj: {} as any,
      response: {} as any,
      isReady: false,
      colors: {
        keyColor: '#ffb939',
        numberColor: '#00d0ff',
        stringColor: '#66da0f',
        trueColor: '#ff7039',
        falseColor: '#ff7039',
        nullColor: '#ff7039',
      },
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

  .body {
    background: #2f2f2f;
    padding: 10px;
    display: flex;

    pre {
      background: #252525;
      padding: 10px;
      border-radius: 4px;
      color: #cecece;
    }
  }
}
</style>
