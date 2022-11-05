<template>
  <section class="share-card open" @click="this.shareCardOperation">
    <h2 class="share-card-title">部屋番号</h2>
    <h2 class="share-card-number">{{ this.$route.params.id }}</h2>
    <div class="share-card-btn" @click="this.copyRoomNumber">
      <v-icon color="orange" size="32" class="share-card-icon">mdi-content-copy</v-icon>
    </div>
  </section>
</template>

<script>
export default {
  methods: {
    shareCardOperation: function () {
      document.querySelector('.share-card').classList.toggle('open');
    },
    copyRoomNumber: function () {
      navigator.clipboard.writeText('https://jphacks-app.herokuapp.com/room/' + this.$route.params.id);
      alert('共有URLがコピーされました!');
    }
  }
};
</script>

<style lang="scss">
.open.share-card {
  position: fixed;
  bottom: 79px;
  right: 0;

  &::before {
    content: '-';
    margin-top: -40px;
    font-size: 56px;
    border-radius: 50%;
  }
}

.share-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: fixed;
  bottom: 36px;
  right: 0;
  padding: 5px 15px;
  padding-right: 45px;
  background-color: #ffdca7;
  border-radius: 24px 0 0 0;
  transition: all 0.4s;

  &::before {
    content: '+';
    color: orange;
    font-size: 36px;
    font-weight: bold;
    line-height: 40px;
    text-align: center;
    position: fixed;
    right: 0;
    margin-top: -80px;
    width: 40px;
    height: 40px;
    background-color: #ffdca7;
    border-radius: 50% 50% 0 0;
    cursor: pointer;
  }

  &-title {
    color: #393939;
    font-size: 16px;
  }

  &-btn {
    &:hover {
      cursor: pointer;
    }
  }

  &-number {
    padding: 0 10px;
    color: #393939;
    font-size: 16px;
  }
}
</style>
