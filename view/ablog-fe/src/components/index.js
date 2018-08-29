import Icon from './Icon.vue';
import Pagination from './Pagination.vue';
import MyButton from './MyButton.vue';
import MyFooter from './MyFooter.vue';
import MyHeader from './MyHeader.vue';
import NavBar from './NavBar.vue';
import ToolBar from './ToolBar.vue';
import SearchBar from './SearchBar.vue';
import EvaluateBox from './EvaluateBox.vue';
import Score from './Score.vue';
import LoadMore from './LoadMore.vue';
import SideBar from './SideBar.vue';
import TocTree from './toctree/TocTree.vue';
import TocTreeItem from './toctree/TocTreeItem.vue';

function install(Vue) {
  Vue.component('icon', Icon);
  Vue.component('pagination', Pagination);
  Vue.component('my-button', MyButton);
  Vue.component('my-footer', MyFooter);
  Vue.component('my-header', MyHeader);
  Vue.component('navbar', NavBar);
  Vue.component('toolbar', ToolBar);
  Vue.component('searchbar', SearchBar);
  Vue.component('evaluate-box', EvaluateBox);
  Vue.component('score', Score);
  Vue.component('load-more', LoadMore);
  Vue.component('sidebar', SideBar);
  Vue.component('toctree-item', TocTreeItem);
  Vue.component('toctree', TocTree);
}

export default {
  install
};
