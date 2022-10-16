<template>
    <div>
        <div>
            <video id="my-video" width="400px" autoplay muted playsinline></video>
        </div>
        <div>
            <video id="their-video" width="400px" autoplay muted playsinline></video>
        </div>
        <h3>peeer id: <span id="my-id"></span></h3>
        <div>
            <v-text-field
            class="d-inline-flex pa-2"
            id="their-id"
            label="peer id"
            :rules="rules"
            hide-details="auto"
            style="width: 50%;"
            ></v-text-field>
            <v-btn id="make-call" color="primary" elevation="4">発信</v-btn>
        </div>
    </div>
</template>

<script>
import Peer from 'skyway-js'

export default {
    layout: "test",
    data() {
        return {
            APIKey: '5152bad7-4798-40b1-986a-a7e8f164a8a3',
            selectedAudio: '',
            selectedVideo: '',
            audios: [],
            videos: [],
            localStream: null,
            peerId: '',
            calltoid: ''
        }
    },

    methods: {
        setEventListener: function(mediaConnection) {
            mediaConnection.on('stream', stream => {
                // video要素にカメラ映像をセットして再生
                const videoElm = document.getElementById('their-video')
                videoElm.srcObject = stream;
                videoElm.play();
            });
        }
    },
    

    mounted: async function() {
        try {
            // カメラ映像取得
            const stream = await navigator.mediaDevices.getUserMedia({video: true, audio: true});
            document.getElementById('my-video').srcObject = stream;
            this.localStream = stream;
        } catch(error) {
            console.log(error);
        }

        //Peer作成
        const peer = new Peer({
            key:   this.APIKey,
            debug: 3,
        });

        //PeerID取得
        peer.on('open', () => {
            document.getElementById('my-id').textContent = peer.id;
        });

        // 発信処理
        document.getElementById('make-call').onclick = () => {
            const theirID = document.getElementById('their-id').value;
            const mediaConnection = peer.call(theirID, this.localStream);
            this.setEventListener(mediaConnection);
        };

        //着信処理(相手から接続要求が来たタイミングで発生)
        peer.on('call', mediaConnection => {
            mediaConnection.answer(this.localStream);
            this.setEventListener(mediaConnection);
        });

        //Peer(相手)との接続が切れた際に発火
        peer.on('close', () => {
            alert('通信が切断しました。');
        });

        //何らかのエラーが発生した場合に発火
        peer.on('error', err => {
            alert(err.message);
        });
    }
};
</script>
