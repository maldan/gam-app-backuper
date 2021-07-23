<template>
  <div class="main">
    <UpdateConfig v-if="showConfig" @close="showConfig = false" />
    <Button @click="showConfig = true" text="Config" icon="add" style="margin-bottom: 10px" />

    <div class="item" v-for="x in list" :key="x">
      <div>{{ x }}</div>
      <div class="amount">{{ backupList[x]?.length }}</div>
      <button @click="backup(x)" class="clickable" :disabled="request[x]">
        BACKUP
        <img src="../asset/preload.svg" alt="" style="width: 16px" />
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Input from '../component/Input.vue';
import Button from '../component/Button.vue';
import UpdateConfig from '../component/UpdateConfig.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: { Input, Button, UpdateConfig },
  async mounted() {
    this.list = await RestApi.main.getList();
    for (let i = 0; i < this.list.length; i++) {
      // @ts-ignore
      this.backupList[this.list[i]] = await RestApi.main.getBackupList(this.list[i]);
    }
  },
  methods: {
    async backup(name: string) {
      // @ts-ignore
      this.request[name] = true;
      try {
        await RestApi.main.backup(name);
      } catch (e) {
        alert(JSON.stringify(e, null, 4));
      }
      // @ts-ignore
      this.request[name] = false;
    },
  },
  data: () => {
    return {
      showConfig: false,
      list: [] as string[],
      request: {},
      backupList: {},
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 10px;

  .item {
    display: flex;
    background: #1d1d1d;
    color: #999999;
    padding: 10px;
    margin-bottom: 10px;
    align-items: center;

    .amount {
      margin-left: auto;
      background: #3c3c3c;
      padding: 2px 15px;
      border-radius: 4px;
    }

    button {
      margin-left: 15px;
      background: #1e44b8;
      color: #b6c8fffe;
      padding: 6px 15px;
      border-radius: 4px;
      border: 0;
      display: flex;
      align-items: center;

      &:disabled {
        background: #4d4f55;
        color: #a1a1a1;
      }

      img {
        display: block;
        margin-left: 10px;
      }
    }
  }
}
</style>
