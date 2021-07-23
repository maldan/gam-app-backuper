<template>
  <div class="container">
    <div class="window">
      <!-- SPACES_KEY -->
      <div class="field">
        <div>SPACES_KEY</div>
        <Input v-model="config.SPACES_KEY" placeholder="SPACES_KEY" />
      </div>

      <!-- SPACES_SECRET -->
      <div class="field">
        <div>SPACES_SECRET</div>
        <Input v-model="config.SPACES_SECRET" placeholder="SPACES_SECRET" />
      </div>

      <!-- SPACES_ENDPOINT -->
      <div class="field">
        <div>SPACES_ENDPOINT</div>
        <Input v-model="config.SPACES_ENDPOINT" placeholder="SPACES_ENDPOINT" />
      </div>

      <!-- SPACES_BUCKET -->
      <div class="field">
        <div>SPACES_BUCKET</div>
        <Input v-model="config.SPACES_BUCKET" placeholder="SPACES_BUCKET" />
      </div>

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" icon="add" />
        <div>&nbsp;</div>
        <Button @click="update()" text="Update" icon="add" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Input from '../component/Input.vue';
import Button from '../component/Button.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: { Input, Button },
  async mounted() {
    this.config = await RestApi.main.getConfig();
  },
  methods: {
    async update() {
      await RestApi.main.updateConfig(this.config);
      this.$emit('close');
    },
  },
  data: () => {
    return {
      config: {
        SPACES_KEY: '',
        SPACES_SECRET: '',
        SPACES_ENDPOINT: '',
        SPACES_BUCKET: '',
      },
    };
  },
});
</script>

<style lang="scss" scoped>
.container {
  padding: 10px;
  background: rgba(0, 0, 0, 0.5);
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .window {
    display: flex;
    flex-direction: column;
    background: #2b2b2b;
    padding: 10px;

    .field {
      margin-bottom: 15px;

      > div {
        color: #999;
        margin-bottom: 10px;
      }
    }
  }
}
</style>
