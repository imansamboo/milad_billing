<?php

namespace App\Controller\Admin;

use EasyCorp\Bundle\EasyAdminBundle\Config\Dashboard;
use EasyCorp\Bundle\EasyAdminBundle\Config\MenuItem;
use EasyCorp\Bundle\EasyAdminBundle\Controller\AbstractDashboardController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;
use App\Entity\Ticket;
use App\Entity\User;
use App\Entity\TicketPost;
use App\Entity\Cdr;
use EasyCorp\Bundle\EasyAdminBundle\Router\CrudUrlGenerator;
use Symfony\Component\Security\Core\User\UserInterface;
use EasyCorp\Bundle\EasyAdminBundle\Config\UserMenu;

class DashboardController extends AbstractDashboardController
{
    /**
     * @Route("/admin", name="admin")
     */
    public function index(): Response
    {
        // redirect to some CRUD controller
        $routeBuilder = $this->get(CrudUrlGenerator::class)->build();

        return $this->redirect($routeBuilder->setController(TicketCrudController::class)->generateUrl());

        // you can also redirect to different pages depending on the current user
//        if ('jane' === $this->getUser()->getUsername()) {
//            return $this->redirect('...');
//        }

        // you can also render some template to display a proper Dashboard
        // (tip: it's easier if your template extends from @EasyAdmin/page/content.html.twig)
//        return $this->render('some/path/my-dashboard.html.twig');
        return parent::index();
    }

    public function configureDashboard(): Dashboard
    {
        return Dashboard::new()
            ->setTitle('Backend');
    }

    public function configureMenuItems(): iterable
    {
        yield MenuItem::linktoDashboard('Dashboard', 'fa fa-home');
        yield MenuItem::section('Tickets');
        yield MenuItem::linkToCrud('Tickets', 'fa fa-ticket', Ticket::class);
        yield MenuItem::linkToCrud('Ticket Posts', 'fa fa-ticket', TicketPost::class);
        yield MenuItem::section('User Management');
        yield MenuItem::linkToCrud('Users', 'fa fa-user', User::class);
        yield MenuItem::section('Cdr');
        yield MenuItem::linkToCrud('Cdr List', 'fa fa-file', Cdr::class);
        yield MenuItem::section('Data Management');

    }
}
