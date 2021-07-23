<template>
  <div :class="$style.select">
    <div v-if="icon" :class="$style.icon">
      <img
        :src="require(`@/asset/${icon}.svg`)"
        alt="Settings"
        draggable="false"
        style="width: 17px"
      />
    </div>

    <input
      type="text"
      :class="$style.input"
      :placeholder="placeholder"
      :value="modelValue ? modelValue.label : ''"
    />

    <button @click="isOpen = !isOpen" :class="$style.icon">
      <img src="../asset/arrow_down.svg" alt="Settings" draggable="false" style="width: 17px" />
    </button>
    <div v-if="isOpen" :class="$style.items">
      <div @click="[select(x), (isOpen = false)]" :class="$style.item" v-for="x in items" :key="x">
        {{ x.label }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    placeholder: String,
    icon: String,
    items: Array,
    modelValue: String,
  },
  async mounted() {},
  methods: {
    select(x: unknown) {
      this.$emit('update:modelValue', x);
    },
  },
  data: () => {
    return {
      isOpen: false,
    };
  },
});
</script>

<style module lang="scss">
.select {
  display: flex;
  align-items: center;
  flex: 1;
  position: relative;

  .items {
    position: absolute;
    left: 0;
    top: 40px;
    width: 100%;
    z-index: 2;

    .item {
      cursor: pointer;
      background: #585858;
      padding: 10px 15px;
      border-bottom: 1px solid #333333;
      color: #9d9d9d;
      user-select: none;
    }
  }

  .icon {
    flex: none;
    background: #515151;
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 4px 0 0 4px;
    border-right: 1px solid #313131;
    box-sizing: border-box;

    img {
      display: block;
    }

    &:last-child {
      border: 0;
      border-radius: 0 4px 4px 0;
      border-left: 1px solid #313131;
    }
  }

  .input {
    box-sizing: border-box;
    background: #3d3d3d;
    color: #9d9d9d;
    border: 0;
    padding: 10px;
    flex: none;
    width: calc(100% - 40px - 40px);
    outline: none;
    font-size: 16px;
    border: 1px solid transparent;

    &:focus {
      border: 1px solid #888888;
    }
  }

  button {
    background: #515151;
    color: #9d9d9d;
    border: 0;
    padding: 10px 15px;
    border-radius: 4px;
    flex: 1;
    font-weight: bold;
    font-size: 16px;
    cursor: pointer;
    user-select: none;
    display: flex;
    align-items: center;

    img {
      display: block;
      margin-left: 15px;
    }

    img:nth-child(1) {
      margin-left: 0;
    }

    &:hover {
      opacity: 0.7;
    }

    &:active {
      position: relative;
      opacity: 0.6;
      top: 1px;
    }
  }
}
</style>
