<script setup lang="ts">
  import { ref } from "vue";
  import { QuillEditor, Delta } from "@vueup/vue-quill";
  import html2canvas from "html2canvas";
  import jsPDF from "jspdf";

  //const pdfContentRef = ref(null);

  const generatePDF = async () => {
    const element: any = document.getElementById("pdf-content");
    const canvas = await html2canvas(element);
    const imgData = canvas.toDataURL("image/png");
    const pdf = new jsPDF("p", "mm", "a4");
    const imgProps = pdf.getImageProperties(imgData);
    const pdfWidth = pdf.internal.pageSize.getWidth();
    const pdfHeight = (imgProps.height * pdfWidth) / imgProps.width;
    pdf.addImage(imgData, "PNG", 0, 0, pdfWidth, pdfHeight);
    pdf.save("download.pdf");
  };
</script>

<template>
  <div id="toolbar-container">
    <span class="ql-formats">
      <select class="ql-font"></select>
      <select class="ql-size"></select>
    </span>
    <span class="ql-formats">
      <button class="ql-bold"></button>
      <button class="ql-italic"></button>
      <button class="ql-underline"></button>
      <button class="ql-strike"></button>
    </span>
    <span class="ql-formats">
      <select class="ql-color"></select>
      <select class="ql-background"></select>
    </span>
    <span class="ql-formats">
      <button class="ql-script" value="sub"></button>
      <button class="ql-script" value="super"></button>
    </span>
    <span class="ql-formats">
      <button class="ql-header" value="1"></button>
      <button class="ql-header" value="2"></button>
      <button class="ql-blockquote"></button>
      <button class="ql-code-block"></button>
    </span>
    <span class="ql-formats">
      <button class="ql-list" value="ordered"></button>
      <button class="ql-list" value="bullet"></button>
      <button class="ql-indent" value="-1"></button>
      <button class="ql-indent" value="+1"></button>
    </span>
    <span class="ql-formats">
      <button class="ql-direction" value="rtl"></button>
      <select class="ql-align"></select>
    </span>
    <span class="ql-formats">
      <button class="ql-link"></button>
      <button class="ql-image"></button>
      <button class="ql-video"></button>
      <button class="ql-formula"></button>
    </span>
    <span class="ql-formats">
      <button class="ql-clean"></button>
    </span>
  </div>

  <button @click="generatePDF">Export PDF</button>

  <QuillEditor id="pdf-content" theme="snow" toolbar="#toolbar-container" />
</template>

<style scoped>
  .ql-container {
    z-index: 1;
    background-color: white;
    margin: auto;
    width: 800px;
    padding: 10px 60px 0;
  }

  .ql-toolbar,
  #toolbar-container {
    z-index: 40;
    position: sticky;
    top: 60px;
    background-color: rgb(247, 247, 247);
    width: 100vw;
    margin: 0;
  }

  #pdf-content {
    width: 210mm; /* A4 width */
    padding: 10mm; /* Optional padding */
  }
</style>
