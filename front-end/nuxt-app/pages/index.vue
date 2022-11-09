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
          <img src="~/assets/img/nomipara_top.png" />
        </div>

        <div class="main-contents-card">
          <h2>部屋を作成する</h2>
          <div class="main-contents-card-sub-msg">
            <p>今すぐオンライン飲み会を始める！🍻</p>
          </div>
          <div class="main-contents-card-btn">
            <Btn text="部屋を作成" color="orange" :clickedfn="this.createRoom" />
          </div>
          <div class="main-contents-card-input-box">
            <h2>部屋番号</h2>
            <div>
              <input type="number" id="room-name" placeholder="部屋番号(例: 1111)" class="main-contents-card-input" />
            </div>
          </div>
          <div>
            <Btn text="部屋に参加" color="blue" :clickedfn="this.joinRoom" />
          </div>
          <p class="main-contents-card-point-ms">
            1ルーム最大18人まで参加できます<br/>
            推奨環境 (PC, Google Chrome)
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';

import Btn from '~/components/presentational/atoms/btn';
import Header from '~/components/presentational/organisms/header';

export default {
  components: {
    Btn,
    Header
  },

  methods: {
    joinRoom: function () {
      const roomName = document.querySelector('#room-name').value;
      this.$router.push('/room/' + roomName);
    },
    createRoom: async function () {
      const response = await axios.post('https://f-2205-server-chhumpv4gq-de.a.run.app/room');
      console.log(response);

      this.$router.push('/room/' + response.data.data.id);
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
      width: 500px;
    }
  }
  &-card {
    box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
    position: relative;
    width: 40%;
    max-width: 900px;
    padding: 35px 15px;
    text-align: center;
    background-color: #fdfdfd;
    color: orange;
    border-radius: 8px;
    & h2 {
      padding: 0 0 15px;
      font-size: 30px;
      font-weight: bold;
    }
    &-sub-msg{
        color: #5a5a5a;
    }
    &-input-box {
      margin: 50px 0 20px;
    }
    &-input {
      border-bottom: solid 2px #ccc;
      background-color: #fff;
      outline: none !important;
      transition: all 0.4s;
      &:hover {
        border-bottom: solid 2px orange;
      }
      &:focus {
        border-bottom: solid 2px orange;
      }
    }
    &-btn {
      margin: 15px 0;
    }
    &-point-ms {
      margin-top: 10px;
      margin-bottom: -20px !important;
      text-align: right;
      font-size: 12px;
      font-weight: normal;
    }
  }
}
</style>
