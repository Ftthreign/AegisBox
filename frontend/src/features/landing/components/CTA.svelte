<script lang="ts">
  import { ArrowRight } from "lucide-svelte";
  import { onMount } from "svelte";

  let container: HTMLDivElement;
  let heading: HTMLHeadingElement;
  let paragraph: HTMLParagraphElement;
  let button: HTMLButtonElement;
  let footerNote: HTMLParagraphElement;

  onMount(async () => {
    const gsap = (await import("gsap")).default;
    const ScrollTrigger = (await import("gsap/ScrollTrigger")).default;

    gsap.registerPlugin(ScrollTrigger);

    // Fade-in utama
    gsap.from(container, {
      opacity: 0,
      y: 40,
      duration: 1,
      ease: "power3.out",
      scrollTrigger: {
        trigger: container,
        start: "top 85%"
      }
    });

    // Stagger elements
    gsap.from([heading, paragraph, button, footerNote], {
      opacity: 0,
      y: 20,
      duration: 0.8,
      stagger: 0.15,
      ease: "power2.out",
      scrollTrigger: {
        trigger: container,
        start: "top 85%"
      }
    });

    // Fade-out saat melewati section
    gsap.to(container, {
      opacity: 0,
      y: 40,
      duration: 0.8,
      ease: "power2.out",
      scrollTrigger: {
        trigger: container,
        start: "bottom top",
        toggleActions: "play none none reverse"
      }
    });
  });

  // Tombol shine effect
  function shine(e: MouseEvent) {
    const target = e.currentTarget as HTMLButtonElement;
    const shineEl = target.querySelector(".shine") as HTMLDivElement;
    shineEl.style.opacity = "1";

    import("gsap").then(({ default: gsap }) => {
      gsap.fromTo(
        shineEl,
        { x: "-120%" },
        { x: "120%", duration: 0.8, ease: "power2.out" }
      );
    });
  }
</script>

<section class="py-24 px-6 bg-gray-50">
  <div
    class="max-w-4xl mx-auto text-center relative"
    bind:this={container}
  >
    <h2
      class="text-4xl md:text-5xl font-bold mb-6 text-gray-900"
      bind:this={heading}
    >
      Ready to Secure Your
      <span class="block text-blue-600">Digital Life?</span>
    </h2>

    <p
      class="text-gray-600 mb-10 max-w-2xl mx-auto leading-relaxed"
      bind:this={paragraph}
    >
      Join thousands of users who trust AegisBox.
    </p>

    <button
      class="relative overflow-hidden px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors flex items-center gap-2 mx-auto"
      bind:this={button}
      on:mouseenter={shine}
    >
      <!-- Shine Layer -->
      <div
        class="shine absolute inset-0 opacity-0 pointer-events-none"
        style="
          background: linear-gradient(
            120deg,
            transparent 0%,
            rgba(255,255,255,0.45) 50%,
            transparent 100%
          );
          transform: translateX(-100%);
        "
      ></div>

      Start Using AegisBox
      <ArrowRight class="w-4 h-4" />
    </button>

    <p class="text-sm text-gray-500 mt-8" bind:this={footerNote}>
      100% Free • Open Source • Self-Hosted
    </p>
  </div>
</section>
