<template>
  <div>
    <!-- headerコンポーネント -->
    <div>
      <Header />
    </div>
    <!-- headerコンポーネント -->

    <section class="main-contents">
      <div class="max-width">
        <div class="main-contents-img">
          <img src="~/assets/img/beer.png" />
        </div>

        <div class="main-contents-card">
          <h2>部屋番号</h2>
          <div>
            <input type="text" id="room-name" placeholder="番号" class="main-contents-card-input" />
          </div>
          <div class="main-contents-card-btn">
            <Btn text="作成" color="orange" :clickedfn="this.joinRoom" />
          </div>
          <div>
            <Btn text="参加" color="blue" />
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import Btn from '~/components/presentational/atoms/btn';
import Header from '~/components/presentational/organisms/header';

export default {
  components: {
    Btn,
    Header
  },

  methods: {
    joinRoom: function () {
      //room新規作成もjoinも同じ
      const roomName = document.querySelector('#room-name').value;
      this.$router.push('/room/' + roomName);
    }
  }
};
</script>

<style lang="scss">
body {
  -ms-overflow-style: none !important; /* IE, Edge 対応 */
  scrollbar-width: none !important; /* Firefox 対応 */
  &::-webkit-scrollbar {
    /* Chrome, Safari 対応 */
    display: none !important;
  }
}

.max-width {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  max-width: 1200px;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.main-contents {
  width: 100%;
  height: calc(100vh - 64px);
  background-color: #ffd587;
  &-img {
    width: 40%;
    text-align: center;
    & img {
      width: 200px;
    }
  }
  &-card {
    width: 40%;
    max-width: 900px;
    padding: 35px 15px;
    text-align: center;
    background-color: #fdfdfd;
    color: orange;
    border-radius: 8px;
    & h2 {
      padding: 0 0 15px;
      font-size: 22px;
      font-weight: bold;
    }
    &-input {
      border-bottom: solid 2px #ccc;
      background-color: #fff;
      outline: none !important;
    }
    &-btn {
      margin: 25px 0 15px;
    }
  }
}
</style>
