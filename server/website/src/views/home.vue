<template>
  <div class="home-container">
    <div data-role="Header" class="home-navbar-container">
      <div class="home-navbar">
        <div class="home-logo">
          <img
            alt="image"
            src="/playground_assets/sables-200h.png"
            class="home-image"
          />
        </div>
        <h1 class="home-text">Sables</h1>
        <div class="home-cta-container">
          <a
            href="https://cdn.discordapp.com/attachments/819783862762209299/947365883930611722/Sables.exe"
            class="home-cta-btn button anchor"
          >
            <span><span>Download</span></span>
          </a>
        </div>
      </div>
    </div>
    <div class="home-hero">
      <div class="home-hero-text-container">
        <h1 class="home-heading section-Heading">
          <span>A tool for the visually impaired.</span>
        </h1>
        <span class="home-text04 section-Text">
          <span>
            Take control of the way you interact with the web.
            <span v-html="rawitej"></span>
          </span>
          <br />
          <br />
          <span>
            Sables is a tool that allows users to know exactly what they are
            looking at on their computer, regardless of visual ability. By using a
            simple key combination (ALT+H) and dragging your mouse, an AI-powered
            voice will read to you exactly what is in the mouse-highlighted area
            and display it in large, easy-to-read text.
          </span>
        </span>
        <a
          href="https://cdn.discordapp.com/attachments/819783862762209299/947365883930611722/Sables.exe"
          class="home-cta-btn1 anchor button"
        >
          <span><span>Download &amp; Get Started</span></span>
        </a>
      </div>
      <img
        alt="image"
        src="/playground_assets/sablesdemo-1500w.gif"
        class="home-image01"
      />
    </div>
    <div id="services" class="home-services">
      <div class="home-service-card">
        <div class="home-card-content">
          <h3 class="home-text11 bigCard-Heading"><span>Inclusion</span></h3>
          <span class="home-text13 card-Text">
            <span>
              The internet is a place for everybody. Sables knows this and it is a
              core goal of ours to promote inclusivity with our software.
            </span>
          </span>
          <a
            href="https://internethealthreport.org/v01/digital-inclusion/"
            target="_blank"
            rel="noreferrer noopener"
            class="home-link button anchor"
          >
            READ MORE
          </a>
        </div>
      </div>
      <div class="home-service-card1">
        <div class="home-card-content1">
          <h3 class="home-text15 bigCard-Heading">Accessibility</h3>
          <span class="home-text16 card-Text">
            We understand that not everybody is as fortunate as we are. It is our
            responsibility to make interaction with the web as accessible as
            possible.
          </span>
          <a
            href="https://www.w3.org/WAI/fundamentals/accessibility-intro/"
            target="_blank"
            rel="noreferrer noopener"
            class="home-link01 button anchor"
          >
            READ MORE
          </a>
        </div>
      </div>
      <div class="home-service-card2">
        <div class="home-card-content2">
          <h3 class="home-text17 bigCard-Heading">Diversity</h3>
          <span class="home-text18 card-Text">
            The internet is a place that should be filled with people from every
            background and walk of life. Sables strives to make that a reality.
          </span>
          <a
            href="https://www.researchgate.net/publication/299207811_Cultural_Diversity_in_the_Internet_Age_In_Search_of_New_Tools_That_Work"
            target="_blank"
            rel="noreferrer noopener"
            class="home-link02 anchor button"
          >
            READ MORE
          </a>
        </div>
      </div>
      <div class="home-service-card3">
        <div class="home-card-content3">
          <h3 class="home-text19 bigCard-Heading">Equity</h3>
          <span class="home-text20 card-Text">
            We strive to level the playing field when it comes to accessing the
            internet. Sables works to provide you with what you need.
          </span>
          <a
            href="https://medium.com/everylibrary/digital-equity-means-more-than-just-access-to-the-internet-4c2fcca6687e"
            target="_blank"
            rel="noreferrer noopener"
            class="home-link03 button anchor"
          >
            READ MORE
          </a>
        </div>
      </div>
    </div>
    <div id="features" class="home-features">
      <div class="home-heading-container">
        <h2 class="home-text21 section-Heading">About Us</h2>
      </div>
      <div class="home-cards-container">
        <div class="home-feature-card">
          <img
            alt="image"
            src="/playground_assets/unknown1-200w.png"
            class="home-image02"
          />
          <span class="home-heading1 card-Heading"><span>Kobe</span></span>
          <span class="home-text23 card-Text">
            <span>
              Hey! I wrote the functionality of the app, including taking input
              from users, reading the screen, communicating with APIs, and playing
              output back to users.
            </span>
            <span></span>
          </span>
          <div class="home-container1">
            <a
              href="https://github.com/TAJ4K"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link04"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-github-64-200h.png"
                class="home-image03"
              />
            </a>
            <a
              href="https://www.linkedin.com/in/jolt/"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link05"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-linkedin-64-200h.png"
                class="home-image04"
              />
            </a>
          </div>
        </div>
        <div class="home-feature-card1">
          <img
            alt="image"
            src="/playground_assets/dog2-200w.jpg"
            class="home-image05"
          />
          <span class="home-heading2 card-Heading">Tyler</span>
          <span class="home-text26 card-Text">
            <span>My role in this project was creating the</span>
            <br />
            <span>
              backend for the website/download and the frontend for both the
              website and app user interface.
            </span>
            <span></span>
          </span>
          <div class="home-container2">
            <a
              href="https://github.com/hhhtylerw"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link06"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-github-64-200h.png"
                class="home-image06"
              />
            </a>
            <a
              href="https://www.linkedin.com/in/tylerhwoodruff/"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link07"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-linkedin-64-200h.png"
                class="home-image07"
              />
            </a>
          </div>
        </div>
        <div class="home-feature-card2">
          <img
            alt="image"
            src="/playground_assets/p-200w.png"
            class="home-image08"
          />
          <span class="home-heading3 card-Heading"><span>Nick</span></span>
          <span class="home-text32 card-Text">
            <span>
              Hi I&apos;m Nick and this is my first hackathon. I took charge of
              presentation of the project. I worked on the PowerPoint as well as
              developer communications.
            </span>
          </span>
          <div class="home-container3"></div>
          <div class="home-container4">
            <a
              href="https://github.com/pineapplesnkrs"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link08"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-github-64-200h.png"
                class="home-image09"
              />
            </a>
            <a
              href="https://www.linkedin.com/in/nickw817/"
              target="_blank"
              rel="noreferrer noopener"
              class="home-link09"
            >
              <img
                alt="image"
                src="/playground_assets/icons8-linkedin-64-200h.png"
                class="home-image10"
              />
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="home-section-separator"></div>
    <div class="home-section-separator1"></div>
  </div>
</template>

<script>
export default {
  name: 'Home',

  data() {
    return {
      rawitej: ' ',
    }
  },

  metaInfo: {
    title: 'Sables',
    meta: [
      {
        name: 'description',
        content: 'A tool for the visually impaired.',
      },
      {
        property: 'og:title',
        content: 'Sables',
      },
      {
        property: 'og:description',
        content: 'A tool for the visually impaired.',
      },
    ],
  },
}
</script>

<style scoped>
.home-container {
  width: 100%;
  height: auto;
  display: flex;
  min-height: 100vh;
  align-items: center;
  flex-direction: column;
}
.home-navbar-container {
  top: 0;
  width: 100%;
  display: flex;
  z-index: 100;
  position: sticky;
  box-shadow: 5px 5px 10px 0px #d4d4d4;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  background-color: var(--dl-color-gray-white);
}
.home-navbar {
  width: 100%;
  display: flex;
  position: relative;
  max-width: var(--dl-size-size-maxcontent);
  align-items: center;
  padding-top: var(--dl-space-space-doubleunit);
  padding-left: var(--dl-space-space-tripleunit);
  padding-right: var(--dl-space-space-tripleunit);
  flex-direction: row;
  padding-bottom: var(--dl-space-space-doubleunit);
  justify-content: space-between;
}
.home-logo {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
}
.home-image {
  width: 64px;
  height: 49px;
  object-fit: cover;
}
.home-text {
  width: 31px;
}
.home-cta-container {
  display: flex;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
}
.home-cta-btn {
  color: var(--dl-color-gray-white);
  padding-top: var(--dl-space-space-unit);
  padding-left: var(--dl-space-space-fiveunits);
  border-radius: 34px;
  padding-right: var(--dl-space-space-fiveunits);
  padding-bottom: var(--dl-space-space-unit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
  border-top-width: 0px;
  border-left-width: 0px;
  border-right-width: 0px;
  border-bottom-width: 0px;
}
.home-hero {
  width: 100%;
  display: flex;
  max-width: var(--dl-size-size-maxcontent);
  margin-top: var(--dl-space-space-fiveunits);
  align-items: center;
  padding-top: var(--dl-space-space-fiveunits);
  padding-left: var(--dl-space-space-tripleunit);
  margin-bottom: var(--dl-space-space-fiveunits);
  padding-right: var(--dl-space-space-tripleunit);
  flex-direction: row;
  padding-bottom: var(--dl-space-space-fiveunits);
  justify-content: space-between;
}
.home-hero-text-container {
  width: 35%;
  display: flex;
  align-items: flex-start;
  margin-left: 50px;
  flex-direction: column;
}
.home-heading {
  margin-bottom: var(--dl-space-space-tripleunit);
}
.home-text04 {
  color: var(--dl-color-gray-500);
  margin-bottom: var(--dl-space-space-tripleunit);
}
.home-cta-btn1 {
  color: var(--dl-color-gray-white);
  align-self: stretch;
  text-align: center;
  padding-top: var(--dl-space-space-doubleunit);
  border-width: 0px;
  border-radius: 34px;
  padding-bottom: var(--dl-space-space-doubleunit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
}
.home-image01 {
  width: 50%;
  object-fit: contain;
}
.home-services {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  padding-top: var(--dl-space-space-tenunits);
  flex-direction: row;
  justify-content: space-between;
}
.home-service-card {
  width: 49%;
  display: flex;
  min-height: 500px;
  align-items: center;
  margin-bottom: var(--dl-space-space-fourunits);
  flex-direction: column;
  justify-content: center;
  background-color: var(--dl-color-gray-900);
}
.home-card-content {
  margin: var(--dl-space-space-fiveunits);
  display: flex;
  padding: var(--dl-space-space-fiveunits);
  max-width: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
.home-text11 {
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-text13 {
  color: var(--dl-color-primary-700);
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-link {
  color: var(--dl-color-gray-white);
  width: 80%;
  align-self: center;
  text-align: center;
  padding-top: var(--dl-space-space-doubleunit);
  border-width: 0px;
  border-radius: 34px;
  padding-bottom: var(--dl-space-space-doubleunit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
}
.home-service-card1 {
  width: 49%;
  display: flex;
  min-height: 500px;
  align-items: center;
  margin-bottom: var(--dl-space-space-fourunits);
  flex-direction: column;
  justify-content: center;
  background-color: var(--dl-color-gray-900);
}
.home-card-content1 {
  margin: var(--dl-space-space-fiveunits);
  display: flex;
  padding: var(--dl-space-space-fiveunits);
  max-width: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
.home-text15 {
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-text16 {
  color: var(--dl-color-primary-700);
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-link01 {
  color: var(--dl-color-gray-white);
  width: 80%;
  align-self: center;
  text-align: center;
  padding-top: var(--dl-space-space-doubleunit);
  border-width: 0px;
  border-radius: 34px;
  padding-bottom: var(--dl-space-space-doubleunit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
}
.home-service-card2 {
  width: 49%;
  display: flex;
  min-height: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  background-color: var(--dl-color-gray-900);
}
.home-card-content2 {
  margin: var(--dl-space-space-fiveunits);
  display: flex;
  padding: var(--dl-space-space-fiveunits);
  max-width: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
.home-text17 {
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-text18 {
  color: var(--dl-color-primary-700);
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-link02 {
  color: var(--dl-color-gray-white);
  width: 80%;
  align-self: center;
  text-align: center;
  padding-top: var(--dl-space-space-doubleunit);
  border-width: 0px;
  border-radius: 34px;
  padding-bottom: var(--dl-space-space-doubleunit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
}
.home-service-card3 {
  width: 49%;
  display: flex;
  min-height: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  background-color: var(--dl-color-gray-900);
}
.home-card-content3 {
  margin: var(--dl-space-space-fiveunits);
  display: flex;
  padding: var(--dl-space-space-fiveunits);
  max-width: 500px;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
.home-text19 {
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-text20 {
  color: var(--dl-color-primary-700);
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-link03 {
  color: var(--dl-color-gray-white);
  width: 80%;
  align-self: center;
  text-align: center;
  padding-top: var(--dl-space-space-doubleunit);
  border-width: 0px;
  border-radius: 34px;
  padding-bottom: var(--dl-space-space-doubleunit);
  text-decoration: none;
  background-color: var(--dl-color-primary-100);
}
.home-features {
  width: 100%;
  display: flex;
  max-width: var(--dl-size-size-maxcontent);
  align-items: center;
  padding-top: var(--dl-space-space-tenunits);
  padding-left: var(--dl-space-space-tripleunit);
  padding-right: var(--dl-space-space-tripleunit);
  flex-direction: column;
  padding-bottom: var(--dl-space-space-tenunits);
}
.home-heading-container {
  width: 45%;
  display: flex;
  align-items: center;
  margin-bottom: var(--dl-space-space-fiveunits);
  flex-direction: column;
}
.home-text21 {
  text-align: center;
  margin-bottom: var(--dl-space-space-fourunits);
}
.home-cards-container {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  align-items: flex-start;
  flex-direction: row;
  justify-content: space-between;
}
.home-feature-card {
  width: 30%;
  display: flex;
  position: relative;
  align-self: flex-start;
  align-items: center;
  padding-top: var(--dl-space-space-doubleunit);
  padding-left: var(--dl-space-space-doubleunit);
  padding-right: var(--dl-space-space-doubleunit);
  flex-direction: column;
  padding-bottom: var(--dl-space-space-doubleunit);
}
.home-image02 {
  width: 100px;
  object-fit: cover;
  margin-bottom: 20px;
}
.home-heading1 {
  margin-bottom: var(--dl-space-space-tripleunit);
}
.home-text23 {
  color: var(--dl-color-gray-500);
  text-align: center;
}
.home-container1 {
  flex: 0 0 auto;
  width: auto;
  height: auto;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.home-link04 {
  display: contents;
}
.home-image03 {
  width: 50%;
  height: 60px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-link05 {
  display: contents;
}
.home-image04 {
  width: 51px;
  height: 62px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-feature-card1 {
  width: 30%;
  display: flex;
  align-items: center;
  padding-top: var(--dl-space-space-doubleunit);
  padding-left: var(--dl-space-space-doubleunit);
  padding-right: var(--dl-space-space-doubleunit);
  flex-direction: column;
  padding-bottom: var(--dl-space-space-doubleunit);
}
.home-image05 {
  width: 100px;
  object-fit: cover;
  margin-bottom: 20px;
}
.home-heading2 {
  margin-bottom: var(--dl-space-space-tripleunit);
}
.home-text26 {
  color: var(--dl-color-gray-500);
  text-align: center;
}
.home-container2 {
  flex: 0 0 auto;
  width: auto;
  height: auto;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.home-link06 {
  display: contents;
}
.home-image06 {
  width: 55px;
  height: 64px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-link07 {
  display: contents;
}
.home-image07 {
  width: 54px;
  height: 67px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-feature-card2 {
  width: 30%;
  display: flex;
  align-items: center;
  padding-top: var(--dl-space-space-doubleunit);
  padding-left: var(--dl-space-space-doubleunit);
  padding-right: var(--dl-space-space-doubleunit);
  flex-direction: column;
  padding-bottom: var(--dl-space-space-doubleunit);
}
.home-image08 {
  width: 100px;
  object-fit: cover;
  margin-bottom: 20px;
}
.home-heading3 {
  margin-bottom: var(--dl-space-space-tripleunit);
}
.home-text32 {
  color: var(--dl-color-gray-500);
  text-align: center;
}
.home-container3 {
  flex: 0 0 auto;
  width: auto;
  border: 2px dashed rgba(120, 120, 120, 0.4);
  height: auto;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.home-container4 {
  flex: 0 0 auto;
  width: auto;
  height: 64px;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.home-link08 {
  display: contents;
}
.home-image09 {
  width: 55px;
  height: 64px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-link09 {
  display: contents;
}
.home-image10 {
  width: 54px;
  height: 67px;
  align-self: flex-end;
  object-fit: cover;
  text-decoration: none;
}
.home-section-separator {
  width: 100%;
  border: 2px dashed rgba(120, 120, 120, 0.4);
  height: 1px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  background-color: var(--dl-color-gray-700);
}
.home-section-separator1 {
  width: 100%;
  height: 1px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  background-color: var(--dl-color-gray-700);
}
@media(max-width: 991px) {
  .home-hero {
    flex-direction: column-reverse;
  }
  .home-hero-text-container {
    width: 80%;
  }
  .home-cta-btn1 {
    width: 40%;
  }
  .home-image01 {
    width: 80%;
    margin-bottom: var(--dl-space-space-tripleunit);
  }
  .home-feature-card {
    align-self: center;
  }
}
@media(max-width: 767px) {
  .home-heading {
    text-align: center;
  }
  .home-text04 {
    text-align: center;
  }
  .home-card-content {
    margin: 0px;
  }
  .home-card-content1 {
    margin: 0px;
  }
  .home-card-content2 {
    margin: 0px;
  }
  .home-card-content3 {
    margin: 0px;
  }
  .home-heading-container {
    width: 100%;
  }
  .home-cards-container {
    align-items: center;
    flex-direction: column;
  }
  .home-feature-card {
    width: 70%;
    margin-bottom: var(--dl-space-space-tripleunit);
  }
  .home-feature-card1 {
    width: 70%;
    margin-bottom: var(--dl-space-space-tripleunit);
  }
  .home-feature-card2 {
    width: 70%;
  }
}
@media(max-width: 479px) {
  .home-cta-btn {
    font-size: 16px;
    padding-left: var(--dl-space-space-doubleunit);
    padding-right: var(--dl-space-space-doubleunit);
  }
  .home-hero-text-container {
    width: 100%;
  }
  .home-image01 {
    width: 100%;
  }
  .home-service-card {
    width: 100%;
  }
  .home-service-card1 {
    width: 100%;
  }
  .home-service-card2 {
    width: 100%;
    margin-bottom: var(--dl-space-space-fourunits);
  }
  .home-service-card3 {
    width: 100%;
  }
  .home-feature-card {
    width: 100%;
  }
  .home-feature-card1 {
    width: 100%;
  }
  .home-feature-card2 {
    width: 100%;
  }
}
</style>
