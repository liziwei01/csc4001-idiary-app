import { Link } from "react-router-dom";

import React, { Component } from "react";

import auth from "./service/authService";

class Landing extends Component {
  state = { user: null };
  componentDidMount() {
    const user = auth.getCurrentUser();
    this.setState({ user });
  }
  render() {
    return (
      <section
        class="landing"
        data-bs-spy="scroll"
        data-bs-target="#navbar"
        data-bs-offset="71"
      >
        <nav
          class="navbar navbar-expand-lg fixed-top navbar-white navbar-custom sticky"
          id="navbar"
        >
          <div class="container">
            <a class="navbar-brand text-uppercase" href="index-1.html">
              <img
                class="logo-light"
                src={require("./img/logo.png")}
                alt=""
                height="50"
              />
            </a>

            <button
              class="navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#navbarCollapse"
              aria-controls="navbarCollapse"
              aria-expanded="false"
              aria-label="Toggle navigation"
            >
              <span class="mdi mdi-menu"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarCollapse">
              <ul class="navbar-nav mx-auto" id="navbar-navlist">
                <li class="nav-item">
                  <a class="nav-link active" href="#home">
                    Home
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#features">
                    Features
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#review">
                    Review
                  </a>
                </li>
              </ul>

              {!this.state.user && (
                <button type="button" class="btn btn-primary nav-btn">
                  <Link className="links text-white" to="./login">
                    Sign Up
                  </Link>
                </button>
              )}
              {this.state.user && (
                <button type="button" class="btn btn-primary nav-btn">
                  <Link className="links text-white" to="./idiary">
                    Welcome to idairy!
                  </Link>
                </button>
              )}
            </div>
          </div>
        </nav>

        <div class="overflow-hidden-x">
          <section class="section home home-6" id="home">
            <div class="bg-overlay"></div>
            <div class="container">
              <div class="row align-items-center">
                <div class="col-lg-7">
                  <div class="home-heading">
                    <h1 class="mb-3 text-white">
                      We're Anxiety Perfect Solution For{" "}
                      <span class="text-warning">Diaries</span>
                    </h1>
                  </div>
                  <p class="fs-20">
                    Explore and learn more about everything to write your diary
                    online and share with more people.
                  </p>
                  <div class="home-btn hstack gap-2 flex-column d-sm-block">
                    <a
                      class="modal-btn"
                      href="javascript:void(0)"
                      data-bs-toggle="modal"
                      data-bs-target=".watchvideomodal"
                    >
                      <span class="avatar-sm">
                        <span class="avatar-title rounded-circle btn-icon">
                          <i class="mdi mdi-play"></i>
                        </span>
                      </span>
                    </a>
                  </div>
                </div>
                <div class="col-lg-5">
                  <div
                    class="mt-5 mt-lg-0"
                    data-aos="fade-down"
                    data-aos-duration="1800"
                  >
                    <img
                      src={require("./img/home-6.png")}
                      class="img-fluid"
                      alt="aaaaaa"
                    />
                    <img
                      class="img-fluid d-none d-lg-block book"
                      src={require("./img/home-6-1.png")}
                      alt="aaaaaaaa"
                    />
                    <img
                      class="img-fluid d-none d-lg-block drive"
                      src={require("./img/home-6-2.png")}
                      alt=""
                    />
                  </div>
                </div>
              </div>
            </div>
          </section>

          <div
            class="modal fade bd-example-modal-lg watchvideomodal"
            data-keyboard="false"
            tabindex="-1"
            aria-hidden="true"
          >
            <div class="modal-dialog modal-dialog-centered modal-dialog modal-lg">
              <div class="modal-content home-modal p-1">
                <div class="modal-header border-0">
                  <button
                    type="button"
                    class="btn-close btn-close-black"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                  ></button>
                </div>
                <video id="VisaChipCardVideo" class="video-box" controls="">
                  <source
                    src="https://www.w3schools.com/html/mov_bbb.mp4"
                    type="video/mp4"
                  />
                </video>
              </div>
            </div>
          </div>

          <section class="section feature" id="features">
            <div class="container">
              <div class="row justify-content-center">
                <div class="col-lg-7">
                  <div class="text-center mb-4">
                    <h2 class="heading">Our Features</h2>
                    <p class="text-muted fs-17">Idiary is a world view</p>
                  </div>
                </div>
              </div>
              <div class="row justify-content-center">
                <div class="col-lg-4 col-md-7 mt-sm-4">
                  <div data-aos="fade-right" data-aos-duration="1800">
                    <div class="feature-card p-3 py-sm-4 rounded d-flex">
                      <div class="flex-shrink-0">
                        <i class="mdi mdi-responsive text-primary float-start me-3 h2"></i>
                      </div>
                      <div class="flex-grow-1 ms-2">
                        <div class="content">
                          <h5 class="title">Private Diary</h5>
                          <p class="text-muted">
                            Always write with words or pictures, and don't miss
                            any moment.
                          </p>
                          <a href="javascipt:void(0)" class="text-dark">
                            Read More <i class="mdi mdi-chevron-right"></i>
                          </a>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col-lg-4 col-md-7 mt-sm-4">
                  <div class="feature-card p-3 py-sm-4 rounded">
                    <i class="mdi mdi-layers-triple-outline text-primary float-start me-3 h2"></i>
                    <div class="content d-block overflow-hidden">
                      <h5 class="title">Personal Homepage</h5>
                      <p class="text-muted mt-2">
                        Social network, a hundred questions and answers about
                        yourself, and all the information is organized in one
                        place.
                      </p>
                      <a href="javascipt:void(0)" class="text-dark">
                        Read More <i class="mdi mdi-chevron-right"></i>
                      </a>
                    </div>
                  </div>
                </div>
                <div class="col-lg-4 col-md-7 mt-sm-4">
                  <div data-aos="fade-left" data-aos-duration="1800">
                    <div class="feature-card p-3 py-sm-4 rounded">
                      <i class="mdi mdi-clipboard-flow-outline text-primary float-start me-3 h2"></i>
                      <div class="content d-block overflow-hidden">
                        <h5 class="title">Explore the World</h5>
                        <p class="text-muted mt-2">
                          Match similar friends through Q and A, or explore more
                          on the world channel.
                        </p>
                        <a href="javascipt:void(0)" class="text-dark">
                          Read More <i class="mdi mdi-chevron-right"></i>
                        </a>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="row my-sm-5 py-5 align-items-center justify-content-between">
                <div class="col-lg-6">
                  <div data-aos="fade-right" data-aos-duration="1800">
                    <div class="card bg-transparent border-0 mb-3 mb-lg-0">
                      <img
                        src={require("./img/feature1.png")}
                        class="img-fluid rounded-3"
                        alt=""
                      />
                    </div>
                  </div>
                </div>
                <div class="col-lg-5">
                  <div data-aos="fade-left" data-aos-duration="1800">
                    <h3 class="feature-heading mb-2">
                      Privacy is our #1 Concern
                    </h3>
                    <p class="text-muted">
                      Even when carefully kept, paper journals can be read by
                      anyone who happens upon them. Penzu keeps your journals
                      safe with double password protection and military strength
                      encryption so you can rest easy knowing that your entries
                      are secure in the Idiary.
                    </p>
                    <ul class="feature-list">
                      <li>
                        <i class="mdi mdi-checkbox-marked-circle-outline text-primary"></i>
                        100% Private
                      </li>
                      <li>
                        <i class="mdi mdi-checkbox-marked-circle-outline text-primary"></i>
                        Available Everywhere
                      </li>
                      <li>
                        <i class="mdi mdi-checkbox-marked-circle-outline text-primary"></i>
                        Never Forget to Write
                      </li>
                    </ul>
                    <a class="btn btn-primary" href="avascipt:void(0)">
                      Learn more
                    </a>
                  </div>
                </div>
              </div>
              <div class="row align-items-center justify-content-between">
                <div class="col-lg-6">
                  <div data-aos="fade-right" data-aos-duration="1800">
                    <h3 class="feature-heading">What are you writing for?</h3>
                    <p class="text-muted">
                      IDiary is your private network. Here, we record our life
                      and exchange opinions, without paying attention to other
                      people's eyes and catering to other people's interests. In
                      the complicated world, iDiary will build a new world for
                      you.
                    </p>
                    <div class="accordion" id="accordionExample">
                      <div class="accordion-item">
                        <h2 class="accordion-header" id="headingOne">
                          <a
                            class="accordion-button"
                            data-bs-toggle="collapse"
                            href="#collapseOne"
                            aria-expanded="true"
                            aria-controls="collapseOne"
                          >
                            <div>
                              <p class="mb-0">Fully Customizable Diary</p>
                              <p class="mb-0 fs-13 text-muted">
                                Our new key fobs make it so easy!
                              </p>
                            </div>
                          </a>
                        </h2>
                        <div
                          id="collapseOne"
                          class="accordion-collapse collapse show"
                          aria-labelledby="headingOne"
                          data-bs-parent="#accordionExample"
                        >
                          <div class="accordion-body">
                            <p class="fs-14 text-muted mb-1">
                              Make each journal your own with custom covers,
                              backgrounds, and fonts.
                            </p>
                            <a href="javascript:void(0)">
                              Check it out <i class="mdi mdi-arrow-right"></i>
                            </a>
                          </div>
                        </div>
                      </div>
                      <div class="accordion-item">
                        <h2 class="accordion-header" id="headingTwo">
                          <a
                            class="accordion-button collapsed"
                            data-bs-toggle="collapse"
                            href="#collapseTwo"
                            aria-expanded="false"
                            aria-controls="collapseTwo"
                          >
                            <div>
                              <p class="mb-0">Smart Journal Search</p>
                              <p class="mb-0 fs-13 text-muted">
                                Keep your engine going with free food and
                                drinks.
                              </p>
                            </div>
                          </a>
                        </h2>
                        <div
                          id="collapseTwo"
                          class="accordion-collapse collapse"
                          aria-labelledby="headingTwo"
                          data-bs-parent="#accordionExample"
                        >
                          <div class="accordion-body">
                            <p class="fs-14 text-muted mb-0">
                              Quickly and easily search through your journals,
                              entries, and tags.
                            </p>
                            <a
                              href="javascript:void(0)"
                              class="fs-14 text-decoration-none"
                            >
                              Check it out <i class="mdi mdi-arrow-right"></i>
                            </a>
                          </div>
                        </div>
                      </div>
                      <div class="accordion-item">
                        <h2 class="accordion-header" id="headingThree">
                          <a
                            class="accordion-button collapsed"
                            data-bs-toggle="collapse"
                            href="#collapseThree"
                            aria-expanded="false"
                            aria-controls="collapseThree"
                          >
                            <div>
                              <p class="mb-0">Military Grade Security</p>
                              <p class="mb-0 fs-13 text-muted">
                                Connectivity, power, and IT issues solved in no
                                time.
                              </p>
                            </div>
                          </a>
                        </h2>
                        <div
                          id="collapseThree"
                          class="accordion-collapse collapse"
                          aria-labelledby="headingThree"
                          data-bs-parent="#accordionExample"
                        >
                          <div class="accordion-body">
                            <p class="fs-14 text-muted mb-0">
                              Further protect your diary with military-grade
                              256-bit AES encryption.
                            </p>
                            <a
                              href="javascript:void(0)"
                              class="fs-14 text-decoration-none"
                            >
                              Check it out <i class="mdi mdi-arrow-right"></i>
                            </a>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col-lg-6">
                  <div data-aos="fade-left" data-aos-duration="1800">
                    <div class="card bg-transparent border-0">
                      <img
                        src={require("./img/feature2.png")}
                        class="img-fluid"
                        alt=""
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <section class="section cta">
            <div class="container">
              <div class="row justify-content-center text-center">
                <div class="col-lg-6">
                  <div data-aos="fade-down" data-aos-duration="1800">
                    <h3 class="fw-bold">Stay Connected</h3>
                    <p>
                      IDiary gives you all the tools you need to focus on the
                      ideas you want to preserve, rather than the process of
                      writing itself.
                    </p>
                    <a class="btn btn-primary" href="javascript:void(0)">
                      Learn More
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <section class="section reviews" id="review">
            <div class="container">
              <div class="row justify-content-center">
                <div class="col-lg-7">
                  <div class="text-center mb-5">
                    <h2 class="heading">
                      What they say about us honest reviews
                    </h2>
                    <p class="text-muted fs-17">
                      Loved by more than 2,000,000 writers, featured on...
                    </p>
                  </div>
                </div>
                <div class="col-lg-8">
                  <div
                    id="carouselExampleIndicators"
                    class="carousel slide"
                    data-bs-ride="carousel"
                  >
                    <div class="carousel-indicators mb-0">
                      <button
                        type="button"
                        data-bs-target="#carouselExampleIndicators"
                        data-bs-slide-to="0"
                        class=""
                        aria-label="Slide 1"
                      ></button>
                      <button
                        type="button"
                        data-bs-target="#carouselExampleIndicators"
                        data-bs-slide-to="1"
                        aria-label="Slide 2"
                        class=""
                      ></button>
                      <button
                        type="button"
                        data-bs-target="#carouselExampleIndicators"
                        data-bs-slide-to="2"
                        aria-label="Slide 3"
                        class="active"
                        aria-current="true"
                      ></button>
                    </div>

                    <div class="carousel-inner">
                      <div class="carousel-item">
                        <div class="card bg-transparent reviews-box h-100">
                          <div class="card-body">
                            <div class="row">
                              <div class="col-md-4">
                                <div class="card reviews-card">
                                  <div class="card-body text-center">
                                    <img
                                      class="img-fluid"
                                      src={require("./img/user/user1.png")}
                                      alt=""
                                    />
                                    <div class="mt-3">
                                      <h6 class="mb-0">Charlotte</h6>
                                      <p class="text-muted mb-0 fs-14">
                                        Actress
                                      </p>
                                    </div>
                                  </div>
                                </div>
                              </div>
                              <div class="col-md-8">
                                <h5 class="reviews-heading">Client Reviews</h5>
                                <div class="d-flex align-items-top">
                                  <div class="flex-shrink-0">
                                    <img
                                      class="reviews-quote-1"
                                      src={require("./img/quote.png")}
                                      alt=""
                                    />
                                  </div>
                                  <div class="flex-grow-1 ms-2 mt-2">
                                    <p class="text-muted fs-14 mt-xl-4">
                                      Originally, I only had the habit of
                                      keeping a diary since I was a child.
                                      Later, I found that the visible range of
                                      my diary was not set, so everyone can see
                                      it. Sometimes it's a negative diary, and
                                      sometimes strangers will encourage
                                      themselves. I think this software is very
                                      warm.
                                    </p>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>

                      <div class="carousel-item active">
                        <div class="card bg-transparent reviews-box h-100">
                          <div class="card-body">
                            <div class="row">
                              <div class="col-md-4">
                                <div class="card reviews-card">
                                  <div class="card-body text-center">
                                    <img
                                      class="img-fluid"
                                      src={require("./img/user/user2.png")}
                                      alt=""
                                    />
                                    <div class="mt-3">
                                      <h6 class="mb-0">Tom</h6>
                                      <p class="text-muted mb-0 fs-14">
                                        Programmer
                                      </p>
                                    </div>
                                  </div>
                                </div>
                              </div>
                              <div class="col-md-8">
                                <h5 class="reviews-heading">Client Reviews</h5>
                                <div class="d-flex align-items-top">
                                  <div class="flex-shrink-0">
                                    <img
                                      class="reviews-quote-1"
                                      src={require("./img/quote.png")}
                                      alt=""
                                    />
                                  </div>
                                  <div class="flex-grow-1 ms-2 mt-2">
                                    <p class="text-muted fs-14 mt-xl-4">
                                      After using the iDiary for so long, talk
                                      about my feeling roughly. The first
                                      impression of iDiary is exquisite. Whether
                                      it is layout or software interaction,
                                      every detail reflects the intention of the
                                      designer.
                                    </p>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>

                      <div class="carousel-item">
                        <div class="card bg-transparent reviews-box h-100">
                          <div class="card-body">
                            <div class="row">
                              <div class="col-md-4">
                                <div class="card reviews-card">
                                  <div class="card-body text-center">
                                    <img
                                      class="img-fluid"
                                      src={require("./img/user/user.png")}
                                      alt=""
                                    />
                                    <div class="mt-3">
                                      <h6 class="mb-0">Selina</h6>
                                      <p class="text-muted mb-0 fs-14">
                                        Writter
                                      </p>
                                    </div>
                                  </div>
                                </div>
                              </div>
                              <div class="col-md-8">
                                <h5 class="reviews-heading">Client Reviews</h5>
                                <div class="d-flex align-items-top">
                                  <div class="flex-shrink-0">
                                    <img
                                      class="reviews-quote-1"
                                      src={require("./img/quote.png")}
                                      alt=""
                                    />
                                  </div>
                                  <div class="flex-grow-1 ms-2 mt-2">
                                    <p class="text-muted fs-14 mt-xl-4">
                                      The most attractive part of iDiary is the
                                      variety of diary recorders. A group of
                                      emotional and loving people. In this field
                                      where no one knows anyone, they expose
                                      their hearts and pry into each other's
                                      privacy.
                                    </p>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <section class="section cta">
            <marquee behavior="" direction="left">
              <p class="mb-0">
                Penzu gives you all the tools you need to focus on the ideas you
                want to preserve, rather than the process of writing itself.
              </p>
            </marquee>
            <div class="container">
              <div class="row justify-content-center text-center">
                <div class="col-lg-6">
                  <div data-aos="fade-down" data-aos-duration="1800">
                    <div class="cta-heading">
                      Today Offer{" "}
                      <span class="mb-3">
                        <span class="counter_value" data-target="37">
                          0
                        </span>
                        <span>% Off</span>
                      </span>
                      !
                    </div>
                    <p>
                      Take your journal wherever you go with our apps for
                      iPhone, iPad, and Android phones and tablets.
                    </p>
                    <p class="fs-18">
                      Limited signup only. Order today before the discount
                      period end.
                    </p>
                    <a class="btn btn-primary" href="javascript:void(0)">
                      Learn More
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <section class="section team">
            <div id="particles-js"></div>

            <div class="container">
              <div class="row justify-content-center">
                <div class="col-lg-7">
                  <div class="text-center mb-5">
                    <h2 class="heading">Team</h2>
                    <p class="text-muted fs-17">
                      Build responsive, mobile-first projects on the web with
                      the world's most popular front-end component library.
                    </p>
                  </div>
                </div>
              </div>
              <div class="row gy-4">
                <div class="col-lg-3 col-sm-6">
                  <div class="team-card">
                    <div class="team-card-img">
                      <img
                        class="img-fluid"
                        src={require("./img/user/user.png")}
                        alt="image"
                      />
                      <div class="team-social-icons">
                        <ul class="list-inline">
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-facebook"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-twitter"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-linkedin"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-instagram"></i>
                            </a>
                          </li>
                        </ul>
                      </div>
                    </div>
                    <div class="team-card-text-2">
                      <h5 class="fw-bold mb-0">Zhang Yiwei</h5>
                      <p class="mb-0 fs-13 text-muted">Front End Developer</p>
                    </div>
                  </div>
                </div>
                <div class="col-lg-3 col-sm-6">
                  <div class="team-card">
                    <div class="team-card-img">
                      <img
                        class="img-fluid"
                        src={require("./img/user/user1.png")}
                        alt="image"
                      />
                      <div class="team-social-icons">
                        <ul class="list-inline">
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-facebook"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-twitter"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-linkedin"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-instagram"></i>
                            </a>
                          </li>
                        </ul>
                      </div>
                    </div>
                    <div class="team-card-text-2">
                      <h5 class="fw-bold mb-0">Yi Kerui</h5>
                      <p class="mb-0 fs-13 text-muted">Front End Developer</p>
                    </div>
                  </div>
                </div>
                <div class="col-lg-3 col-sm-6">
                  <div class="team-card">
                    <div class="team-card-img">
                      <img
                        class="img-fluid"
                        src={require("./img/user/user2.png")}
                        alt="image"
                      />
                      <div class="team-social-icons">
                        <ul class="list-inline">
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-facebook"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-twitter"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-linkedin"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-instagram"></i>
                            </a>
                          </li>
                        </ul>
                      </div>
                    </div>
                    <div class="team-card-text-2">
                      <h5 class="fw-bold mb-0">Li Ziwei</h5>
                      <p class="mb-0 fs-13 text-muted">Back End Developer</p>
                    </div>
                  </div>
                </div>
                <div class="col-lg-3 col-sm-6">
                  <div class="team-card">
                    <div class="team-card-img">
                      <img
                        class="img-fluid"
                        src={require("./img/user/user3.png")}
                        alt="image"
                      />
                      <div class="team-social-icons">
                        <ul class="list-inline">
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-facebook"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-twitter"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-linkedin"></i>
                            </a>
                          </li>
                          <li class="list-inline-item">
                            <a href="javascript:void(0)" target="">
                              <i class="mdi mdi-instagram"></i>
                            </a>
                          </li>
                        </ul>
                      </div>
                    </div>
                    <div class="team-card-text-2">
                      <h5 class="fw-bold mb-0">Wu Kexin</h5>
                      <p class="mb-0 fs-13 text-muted">Back End Developer</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <footer class="section footer">
            <div class="container">
              <div class="row justify-content-center">
                <div class="col-lg-6 col-sm-10 text-center">
                  <a href="javascript:void(0)">
                    <img src={require("./img/logo.png")} height="50" alt="" />
                  </a>
                  <p class="mx-auto mt-sm-4">
                    Obviously I'm a Web Designer. Experienced with all stages of
                    the development cycle for dynamic web projects.
                  </p>
                  <ul class="list-unstyled mb-0 mt-4 social-icon">
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-facebook"></i>
                      </a>
                    </li>
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-twitter"></i>
                      </a>
                    </li>
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-instagram"></i>
                      </a>
                    </li>
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-vimeo"></i>
                      </a>
                    </li>
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-google-plus"></i>
                      </a>
                    </li>
                    <li class="list-inline-item">
                      <a href="javascript:void(0)">
                        <i class="mdi mdi-linkedin"></i>
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
              <div class="row justify-content-center">
                <div class="col-lg-12">
                  <div class="text-center mt-3">
                    <ul class="list-unstyled mb-0">
                      <li class="list-inline-item mx-lg-3 m-2">
                        <a class="text-white" href="javascript:void(0)">
                          Home
                        </a>
                      </li>
                      <li class="list-inline-item mx-lg-3 m-2">
                        <a class="text-white" href="javascript:void(0)">
                          About us
                        </a>
                      </li>
                      <li class="list-inline-item mx-lg-3 m-2">
                        <a class="text-white" href="javascript:void(0)">
                          Order
                        </a>
                      </li>
                      <li class="list-inline-item mx-lg-3 m-2">
                        <a class="text-white" href="javascript:void(0)">
                          Member
                        </a>
                      </li>
                      <li class="list-inline-item mx-lg-3 m-2">
                        <a class="text-white" href="javascript:void(0)">
                          Contact Us
                        </a>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </footer>

          <div class="footer-alt pt-3 pb-3">
            <div class="container">
              <div class="row">
                <div class="col-lg-12">
                  <div class="text-center">
                    <p class="mb-0 text-white fs-15">
                      ©<script>document.write(new Date().getFullYear())</script>{" "}
                      Idiary. Design by
                      <i class="mdi mdi-heart text-danger"></i> Themesdesign
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    );
  }
}

export default Landing;
