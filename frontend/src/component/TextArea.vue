<template>
  <div :class="$style.textarea">
    <textarea :placeholder="placeholder" :value="modelValue" @input="change"></textarea>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    placeholder: String,
    modelValue: String,
  },
  async mounted() {},
  methods: {
    change(e: any) {
      this.$emit('update:modelValue', e.target.value);
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style module lang="scss">
.textarea {
  display: flex;
  align-items: center;
  flex: 1;
  box-sizing: border-box;

  textarea {
    box-sizing: border-box;
    background: #3d3d3d;
    color: #9d9d9d;
    border: 0;
    padding: 10px;
    flex: 1;
    outline: none;
    font-size: 16px;
    border: 1px solid transparent;

    min-width: 0;

    &:focus {
      border: 1px solid #888888;
    }
  }
}
</style>
