<template>
  <section class="loader">
    <div class="loader-back">
      <div class="loader-back-wave-box">
        <svg viewBox="0 0 2700 300" width="2700" height="300">
          <path
            d="
            M 0,0 
            v 100,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            t 300,0 
            q 150,50 300,0 
            v 0,-100 
            Z
            "
          />
        </svg>
      </div>
    </div>
    <div class="loader-front">
      <img src="~/assets/img/loaderImg.svg" class="loader-front-img" />
      <h2 class="loader-front-ms">Loading Now...</h2>
    </div>
  </section>
</template>

<script>
export default {
  props: ['autoDelete'],
  mounted: function () {
    if (!this.autoDelete) return;

    setTimeout(() => {
      this.$emit('loaderOperationFn');
    }, 4000);
  }
};
</script>

<style lang="scss">
.loader {
  position: fixed;
  top: 0;
  left: 0;

  &-back {
    width: 100vw;
    height: 100vh;
    background-color: orange;
    z-index: 1000;
    overflow: hidden;

    &-wave-box {
      height: 82vh;
      background-color: #ffd89c;
      animation: wave-box linear 3s forwards;
      & svg {
        transform: translateY(80vh);
        animation: wave-box-svg linear 3s forwards;
        & path {
          fill: #ffd89c;
          animation: wave linear 6s infinite;
        }
      }
    }
  }
  &-front {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    &-ms {
      padding: 50px 0;
      text-align: center;
      color: #474747;
      font-size: 34px;
    }

    &-img {
      width: 300px;
    }
  }
}

@keyframes wave {
  from {
    transform: translate(0);
  }
  to {
    transform: translate(-66.66%);
  }
}

@keyframes wave-box {
  from {
    height: 82vh;
  }
  to {
    height: 0px;
  }
}

@keyframes wave-box-svg {
  from {
    transform: translateY(80vh);
  }
  to {
    transform: translateY(-50px);
  }
}
</style>
